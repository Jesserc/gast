package transaction

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// createRawTransaction creates a raw Ethereum transaction.
func createRawTransaction(rpcURL, to, data, privateKey string, gasLimit, wei uint64) (string, error) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	// Get chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return "", err
	}

	// Get base fee
	baseFee, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Println("base fee:", baseFee)

	// Get gas tip cap
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}

	// Calculate gas fee cap with 2 Gwei margin
	increment := big.NewInt(1e9)
	gasFeeCap := new(big.Int).Add(gasTipCap, increment)

	fmt.Println("max fee: per gas", gasFeeCap)

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
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Convert data to hex format
	hexData := "0x" + hex.EncodeToString([]byte(data))
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
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     amount,
		Data:      bytesData,
	}

	tx := types.NewTx(&txData)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		fmt.Println("line 82:", err)
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
