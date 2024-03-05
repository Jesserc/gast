package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
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

func CreateContract(rpcURL, data, privateKey string, gasLimit, wei uint64) TxReceipt {
	var receipt TxReceipt

	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	defer client.Close()
	if err != nil {
		log.Crit("Failed to dial RPC client", "error", err)
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Crit("Failed to get chain ID", "error", err)
	}

	// Get base fee
	baseFee, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Crit("Failed to get base fee", "error", err)
	}
	log.Info("base fee", "value", baseFee)

	// Get priority fee
	priorityFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Crit("Failed to get priority fee", "error", err)
	}
	log.Info("priority fee", "value", priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, priorityFee)

	log.Info("max fee per gas", "value", gasFeeCap)

	// Decode private key
	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		log.Crit("Failed to decode private key", "error", err)
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		log.Crit("Failed to convert private key to ECDSA", "error", err)
	}

	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Crit("Failed to get nonce", "error", err)
	}

	data = strings.Trim(data, "\n")

	// Add 0x prefix
	data = "0x" + data

	bytesData, err := hexutil.Decode(data)
	if err != nil {
		log.Crit("Failed to decode data", "error", err)
	}

	// Create transaction data
	amount := new(big.Int).SetUint64(wei)

	tx, err := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		Value:     amount,
		Data:      bytesData,
	}), nil
	if err != nil {
		log.Crit("Failed to create transaction", "error", err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		log.Crit("Failed to sign transaction", "error", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	timer := time.Now()

	fmt.Println() // spacing
	log.Warn("Sending transaction, please wait for confirmation...")
	
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Crit("Failed to send transaction", "error", err)
	}

	var transactionURL string
	for id, explorer := range gastParams.NetworkExplorers {
		if chainID.Uint64() == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, signedTx.Hash().Hex())
			break
		}
	}

	select {
	case <-time.After(35 * time.Second):
		log.Crit("Timeout:", "time taken", time.Since(timer))
	case <-ctx.Done():
		_, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		if err != nil {
			log.Crit("Failed to get transaction status", "error", err)
		}

		if isPending {
			log.Info("Transaction update", "", "Transaction is still pending")
			fmt.Println("Tx receipt:", transactionURL)
			os.Exit(0)
		} else {
			tr, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
			if err != nil {
				log.Crit("Failed to get transaction receipt", "error", err)
			}

			trBytes, err := tr.MarshalJSON()
			if err != nil {
				log.Crit("Failed to marshal transaction receipt", "error", err)
			}

			err = receipt.UnmarshalJSON(trBytes)
			if err != nil {
				log.Crit("Failed to unmarshal transaction bytes to type Go TxReceipt", "error", err)
			}
			receipt.TransactionURL = transactionURL
		}
	}

	return receipt
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
