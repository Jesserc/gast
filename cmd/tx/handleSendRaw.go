package transaction

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/Jesserc/gast/cmd/tx/params"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

// networkExplorers maps network IDs to their respective explorers.
var networkExplorers = map[uint64]string{
	0x01:     "https://etherscan.io/",                          // Ethereum Mainnet
	0x05:     "https://goerli.etherscan.io/",                   // Goerli Testnet
	0xAA36A7: "https://sepolia.etherscan.io/",                  // Sepolia Testnet
	0x89:     "https://polygonscan.com/",                       // Polygon Mainnet
	0x13881:  "https://mumbai.polygonscan.com/",                // Polygon Mumbai Testnet
	0x0A:     "https://optimistic.etherscan.io/",               // Optimism Mainnet
	0x1A4:    "https://goerli-optimism.etherscan.io/",          // Optimism Goerli Testnet
	0xA4B1:   "https://arbiscan.io/",                           // Arbitrum One Mainnet
	0x66EEE:  "https://sepolia.arbiscan.io/",                   // Arbitrum Sepolia Testnet
	0x38:     "https://bscscan.com/",                           // Binance Smart Chain Mainnet
	0x61:     "https://testnet.bscscan.com/",                   // Binance Smart Chain Testnet
	0x421611: "https://explorer.celo.org/",                     // Celo Mainnet
	0xA4EC:   "https://alfajores-blockscout.celo-testnet.org/", // Celo Alfajores Testnet
	0x2105:   "https://basescan.org/",                          // Base Mainnet
	0xE708:   "https://lineascan.build/",                       // Linea Mainnet
	0x144:    "https://explorer.zksync.io/",                    // zkSync Mainnet
}

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

// sendRawTransaction sends a raw Ethereum transaction.
func sendRawTransaction(rawTx, rpcURL string) (string, error) {
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return "", err
	}

	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	var transactionURL string
	for id, explorer := range networkExplorers {
		if chainID.Uint64() == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, tx.Hash().Hex())
		}
	}

	// Unmarshal the transaction JSON into a struct
	var txDetails Transaction
	txBytes, err := tx.MarshalJSON()
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(txBytes, &txDetails); err != nil {
		return "", err
	}

	// Add additional transaction details
	txDetails.TransactionTime = tx.Time().Format(time.RFC822)
	txDetails.TransactionCost = tx.Cost().String()

	// Convert some hexadecimal string fields to decimal string
	convertFields := []string{"Nonce", "MaxPriorityFeePerGas", "MaxFeePerGas", "Value", "Type", "Gas"}
	for _, field := range convertFields {
		if err := convertHexField(&txDetails, field); err != nil {
			return "", err
		}
	}

	// Marshal the struct back to JSON
	txJSON, err := json.MarshalIndent(txDetails, "", "\t")
	if err != nil {
		return "", err
	}

	// Print the entire JSON with the added fields
	fmt.Println(params.ColorGreen, "Transaction details:", params.ColorReset)
	fmt.Println(string(txJSON))
	
	transactionReceipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txDetails.Hash))
	if err != nil {
		return "", err
	}
	fmt.Println(params.ColorGreen, "Transaction receipt:", params.ColorReset)
	fmt.Println(transactionReceipt)

	return transactionURL, nil
}

func convertHexField(tx *Transaction, field string) error {
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

type Tx struct {
	Type                 string        `json:"type"`
	ChainId              string        `json:"chainId"`
	Nonce                string        `json:"nonce"`
	To                   string        `json:"to"`
	Gas                  string        `json:"gas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	Value                string        `json:"value"`
	Input                string        `json:"input"`
	AccessList           []interface{} `json:"accessList"`
	V                    string        `json:"v"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	YParity              string        `json:"yParity"`
	Hash                 string        `json:"hash"`
	TransactionTime      string        `json:"transactionTime"`
	TransactionCost      string        `json:"transactionCost"`
}
