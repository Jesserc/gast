package transaction

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Jesserc/gast/cmd/tx/gastParams"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
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

func CreateContract(rpcURL, data, privateKey string, gasLimit, wei uint64) (TxReceipt, error) {
	var receipt TxReceipt

	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	defer client.Close()
	if err != nil {
		return TxReceipt{}, err
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return TxReceipt{}, err
	}

	// Get base fee
	baseFee, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return TxReceipt{}, err
	}
	fmt.Printf("%sbase fee:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, baseFee)

	// Get priority fee
	priorityFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return TxReceipt{}, err
	}
	fmt.Printf("%spriority fee:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, priorityFee) /* .Add(increment, big.NewInt(0)) */

	fmt.Printf("%smax fee per gas:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, gasFeeCap)

	// Decode private key
	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return TxReceipt{}, err
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		return TxReceipt{}, err
	}

	publicKey := ecdsaPrivateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err := fmt.Errorf("error casting public key to ECDSA")
		return TxReceipt{}, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return TxReceipt{}, err
	}

	strings.Trim(data, "\n")
	strings.Trim(data, "\n")

	// Add 0x prefix
	data = "0x" + data

	bytesData, err := hexutil.Decode(data)
	if err != nil {
		return TxReceipt{}, err
	}

	// Create transaction data
	amount := new(big.Int).SetUint64(wei)

	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		Value:     amount,
		Data:      bytesData,
	}

	tx := types.NewTx(&txData)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		return TxReceipt{}, err
	}

	// Encode signed transaction to RLP hex
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)
	if err != nil {
		return TxReceipt{}, err
	}

	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	rawTxBytes, err := hex.DecodeString(rawTxRLPHex)
	if err != nil {
		return TxReceipt{}, err
	}

	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return TxReceipt{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		return TxReceipt{}, err
	}

	var transactionURL string
	for id, explorer := range gastParams.NetworkExplorers {
		if chainID.Uint64() == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, tx.Hash().Hex())
			break
		}
	}

	select {
	case val := <-time.After(30 * time.Second):
		fmt.Println("timeout:", val.Format(time.Kitchen))
		fmt.Println("tx receipt:", transactionURL)
		os.Exit(0)
	case <-ctx.Done():
		_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			return TxReceipt{}, err
		}

		if isPending {
			fmt.Println("transaction is still pending")
			fmt.Println("tx receipt:", transactionURL)
			os.Exit(0)
		} else {
			tr, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				return TxReceipt{}, err
			}

			trBytes, err := tr.MarshalJSON()
			if err != nil {
				return TxReceipt{}, err
			}

			err = receipt.UnmarshalJSON(trBytes)
			if err != nil {
				return TxReceipt{}, err
			}
			receipt.TransactionURL = transactionURL
		}
	}

	return receipt, nil
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
