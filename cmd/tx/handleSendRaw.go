package transaction

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction represents the structure of the transaction JSON.
type Transaction struct {
	Type                 string   `json:"type"`
	ChainID              string   `json:"chainId"`
	Nonce                string   `json:"nonce"`
	To                   string   `json:"to"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice,omitempty"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	Value                string   `json:"value"`
	Input                string   `json:"input"`
	AccessList           []string `json:"accessList"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	YParity              string   `json:"yParity"`
	Hash                 string   `json:"hash"`
	TransactionTime      string   `json:"transactionTime,omitempty"`
	TransactionCost      string   `json:"transactionCost,omitempty"`
}

// SendRawTransaction sends a raw Ethereum transaction.
func SendRawTransaction(rawTx, rpcURL string) (string, string) {
	client, err := ethclient.Dial(rpcURL)
	defer client.Close()
	if err != nil {
		log.Crit("Failed to dial RPC client", "error", err)
	}

	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Crit("Failed to decode raw transaction to rlp decoded bytes", "error", err)
	}

	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		log.Crit("Failed to decode transaction rlp bytes to types.Transaction", "error", err)
	}

	fmt.Println() // spacing
	log.Warn("Sending transaction, please wait for confirmation...")

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Crit("Failed to send transaction", "error", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Crit("Failed to get chain ID", "error", err)
	}

	var transactionURL string
	for id, explorer := range gastParams.NetworkExplorers {
		if chainID.Uint64() == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, tx.Hash().Hex())
			break
		}
	}

	// Unmarshal the transaction JSON into a struct
	var txDetails Transaction
	txBytes, err := tx.MarshalJSON()
	if err != nil {
		log.Crit("Failed to marshal transaction", "error", err)
	}
	if err := json.Unmarshal(txBytes, &txDetails); err != nil {
		log.Crit("Failed to unmarshal transaction bytes to type Go Transaction", "error", err)
	}

	// Add additional transaction details
	txDetails.TransactionTime = tx.Time().Format(time.RFC822)
	txDetails.TransactionCost = tx.Cost().String()

	// Convert some hexadecimal string fields to decimal string
	convertFields := []string{"Nonce", "MaxPriorityFeePerGas", "MaxFeePerGas", "Value", "Type", "Gas"}
	for _, field := range convertFields {
		if err := convertHexFieldToDecimalString(&txDetails, field); err != nil {
			log.Error("Failed to convert hex fields to decimal string", "error", err)
		}
	}

	// Marshal the struct back to JSON
	txJSON, err := json.MarshalIndent(txDetails, "", "\t")
	if err != nil {
		log.Crit("Failed to marshal indent transaction details", "error", err)
	}

	return transactionURL, string(txJSON)
}

func convertHexFieldToDecimalString(tx *Transaction, field string) error {
	// Get the type of the Transaction struct
	typeOfTx := reflect.TypeOf(*tx)

	// Get the value of the Transaction struct
	txValue := reflect.ValueOf(tx).Elem()

	// Parse the hexadecimal string as an integer
	hexStr := txValue.FieldByName(field).String()

	intValue, err := strconv.ParseUint(hexStr[2:], 16, 64)
	if err != nil {
		return err
	}

	// Convert the integer to a decimal string
	decimalStr := strconv.FormatUint(intValue, 10)

	// Check if the field exists
	_, ok := typeOfTx.FieldByName(field)
	if !ok {
		return fmt.Errorf("field %s does not exist in Transaction struct", field)
	}

	// Set the field value to the decimal string
	txValue.FieldByName(field).SetString(decimalStr)

	return nil
}
