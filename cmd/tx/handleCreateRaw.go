package transaction

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/Jesserc/gast/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

// CreateRawTransaction creates a raw Ethereum transaction.
func CreateRawTransaction(rpcURL, to, data, privateKey string, gasLimit, wei uint64) (string, error) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	defer client.Close()
	if err != nil {
		return "", err
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	// Get base fee
	baseFee, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Printf("%sbase fee:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, baseFee)

	// Get priority fee
	priorityFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Printf("%spriority fee:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, priorityFee) /* .Add(increment, big.NewInt(0)) */

	fmt.Printf("%smax fee per gas:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, gasFeeCap)

	// Decode private key
	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return "", err
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		return "", err
	}

	publicKey := ecdsaPrivateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA pubKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Convert data to hex format
	var hexData string
	if !utils.IsHexWithOrWithout0xPrefix(data) {
		hexData = hexutil.Encode([]byte(data))
	} else if strings.HasPrefix(data, "0x") {
		hexData = data
	} else {
		hexData = "0x" + data
	}

	bytesData, err := hexutil.Decode(hexData)
	if err != nil {
		return "", err
	}

	// Create transaction data
	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(wei)
	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     amount,
		Data:      bytesData,
	}

	tx := types.NewTx(&txData)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	// Encode signed transaction to RLP hex
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)
	if err != nil {
		return "", err
	}

	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	return rawTxRLPHex, nil
}
