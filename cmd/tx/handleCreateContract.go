package transaction

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Jesserc/gast/cmd/gastParams"
	rpcfactory "github.com/Jesserc/gast/internal/rpc_factory"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/w3"
	w3eth "github.com/lmittmann/w3/module/eth"
)

type TxReceipt struct {
	Type              uint64        `json:"type"`
	Root              string        `json:"root"`
	Status            uint64        `json:"status"`
	CumulativeGasUsed uint64        `json:"cumulativeGasUsed"`
	LogsBloom         string        `json:"logsBloom,omitempty"`
	Logs              []interface{} `json:"logs,omitempty"`
	TransactionHash   string        `json:"transactionHash"`
	ContractAddress   string        `json:"contractAddress,omitempty"`
	GasUsed           uint64        `json:"gasUsed"`
	EffectiveGasPrice uint64        `json:"effectiveGasPrice"`
	TransactionCost   uint64        `json:"-"`
	BlockHash         string        `json:"blockHash"`
	BlockNumber       uint64        `json:"blockNumber"`
	TransactionIndex  uint64        `json:"transactionIndex"`
	TransactionURL    string        `json:"-"`
}

func CreateContract(rpcURL, data, privateKey string, gasLimit, wei uint64) (*TxReceipt, error) {
	var receipt TxReceipt

	// Connect to the Ethereum client
	client, err := w3.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC client: %s", err)
	}
	defer client.Close()

	var (
		chainID     uint64
		baseFee     big.Int
		priorityFee big.Int
		errs        w3.CallErrors
	)

	if err := client.CallCtx(
		context.Background(),
		w3eth.ChainID().Returns(&chainID),
		w3eth.GasPrice().Returns(&baseFee),
		w3eth.GasTipCap().Returns(&priorityFee),
	); errors.As(err, &errs) {
		if errs[0] != nil {
			return nil, fmt.Errorf("failed to get chain: %s", err)
		} else if errs[1] != nil {
			return nil, fmt.Errorf("failed to get base fee: %s", err)
		} else if errs[2] != nil {
			return nil, fmt.Errorf("failed to get priority fee: %s", err)
		}
	}
	log.Info("base fee", "value", baseFee)
	log.Info("priority fee", "value", priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(&baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, &priorityFee)

	log.Info("max fee per gas", "value", gasFeeCap)

	// Decode private key
	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %s", err)
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert private key to ECDSA: %s", err)
	}

	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)

	var pendingNonce string
	if err := client.CallCtx(
		context.Background(),
		rpcfactory.PendingNonceAt(fromAddress).Returns(&pendingNonce),
	); err != nil {
		return nil, fmt.Errorf("failed to get nonce: %s", err)
	}

	nonce, err := hexutil.DecodeUint64(pendingNonce)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nonce: %s", err)
	}

	data = strings.Trim(data, "\n")

	// Add 0x prefix
	data = "0x" + data

	bytesData, err := hexutil.Decode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %s", err)
	}

	// Create transaction data
	amount := new(big.Int).SetUint64(wei)

	tx, err := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonce,
		GasTipCap: &priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		Value:     amount,
		Data:      bytesData,
	}), nil
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %s", err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(int64(chainID))), ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	timer := time.Now()

	fmt.Println() // spacing
	log.Warn("Sending transaction, please wait for confirmation...")

	var hash common.Hash
	if err := client.CallCtx(
		ctx,
		w3eth.SendTx(signedTx).Returns(&hash),
	); err != nil {
		return nil, fmt.Errorf("failed to send transaction: %s", err)
	}

	var transactionURL string
	for id, explorer := range gastParams.NetworkExplorers {
		if chainID == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, signedTx.Hash().Hex())
			break
		}
	}

	select {
	case <-time.After(35 * time.Second):
		log.Crit("Timeout:", "time taken", time.Since(timer))
	case <-ctx.Done():
		var txn types.Transaction
		if err := client.CallCtx(
			context.Background(),
			w3eth.Tx(signedTx.Hash()).Returns(&txn),
		); err != nil {
			return nil, fmt.Errorf("failed to get transaction status: %s", err)
		}

		isPending := true
		if _, r, _ := txn.RawSignatureValues(); r != nil {
			isPending = false
		}

		if isPending {
			log.Info("Transaction update", "", "Transaction is still pending")
			fmt.Println("Tx receipt:", transactionURL)
			return nil, nil
		} else {
			var tr types.Receipt
			if err := client.CallCtx(
				context.Background(),
				w3eth.TxReceipt(signedTx.Hash()).Returns(&tr),
			); err != nil {
				return nil, fmt.Errorf("failed to get transaction receipt: %s", err)
			}

			trBytes, err := tr.MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed to marshal transaction receipt: %s", err)
			}

			err = receipt.UnmarshalJSON(trBytes)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal transaction bytes to type Go TxReceipt: %s", err)
			}
			receipt.TransactionURL = transactionURL
		}
	}

	return &receipt, nil
}

// UnmarshalJSON customizes the unmarshalling of a TxReceipt.
func (r *TxReceipt) UnmarshalJSON(data []byte) error {
	// Define a helper struct inside the method with the fields as strings
	type Alias TxReceipt
	helper := struct {
		Type              string `json:"type"`
		Status            string `json:"status"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		GasUsed           string `json:"gasUsed"`
		EffectiveGasPrice string `json:"effectiveGasPrice"`
		BlockNumber       string `json:"blockNumber"`
		TransactionIndex  string `json:"transactionIndex"`
		*Alias
	}{
		Alias: (*Alias)(r), // Point Alias to the TxReceipt's memory
	}

	if err := json.Unmarshal(data, &helper); err != nil {
		return err
	}

	// Convert hexadecimal fields to uint64
	var err error
	if r.Type, err = hexToUint64(helper.Type); err != nil {
		return fmt.Errorf("error parsing Type: %v", err)
	}
	if r.Status, err = hexToUint64(helper.Status); err != nil {
		return fmt.Errorf("error parsing Status: %v", err)
	}
	if r.CumulativeGasUsed, err = hexToUint64(helper.CumulativeGasUsed); err != nil {
		return fmt.Errorf("error parsing CumulativeGasUsed: %v", err)
	}
	if r.GasUsed, err = hexToUint64(helper.GasUsed); err != nil {
		return fmt.Errorf("error parsing GasUsed: %v", err)
	}
	if r.EffectiveGasPrice, err = hexToUint64(helper.EffectiveGasPrice); err != nil {
		return fmt.Errorf("error parsing EffectiveGasPrice: %v", err)
	}
	if r.BlockNumber, err = hexToUint64(helper.BlockNumber); err != nil {
		return fmt.Errorf("error parsing BlockNumber: %v", err)
	}
	if r.TransactionIndex, err = hexToUint64(helper.TransactionIndex); err != nil {
		return fmt.Errorf("error parsing TransactionIndex: %v", err)
	}

	r.TransactionCost = r.EffectiveGasPrice * r.GasUsed

	return nil
}

// hexToUint64 converts a hexadecimal string to uint64.
func hexToUint64(hexStr string) (uint64, error) {
	return strconv.ParseUint(hexStr, 0, 64)
}

func CompileSol(dir string) (string, error) {
	var bytecode string
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("solc", dir, "--bin")
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	output := stdoutBuf.String()

	regex := regexp.MustCompile(`6080[0-9a-fA-F]+`)
	matches := regex.FindStringSubmatch(output)
	if len(matches) > 0 {
		bytecode = matches[0]
	} else {
		err = fmt.Errorf("bytecode not found")
		return "", err
	}

	return bytecode, nil
}
