package transaction

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	gastParam "github.com/Jesserc/gast/cmd/tx/gastParams"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

// createRawTransaction creates a raw Ethereum transaction.
func createRawTransaction(rpcURL, to, data, privateKey string, gasLimit, wei uint64) (string, error) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	fmt.Println(len([]byte("0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5")))
	pubkey, err := crypto.DecompressPubkey([]byte("0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pubkey)
	common.HexToAddress("0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5")
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
	fmt.Printf("%sbase fee:%s %s\n", gastParam.ColorGreen, gastParam.ColorReset, baseFee)

	// Get priority fee
	priorityFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Printf("%spriority fee:%s %s\n", gastParam.ColorGreen, gastParam.ColorReset, priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(baseFee, big.NewInt(2*params.GWei))
	fmt.Println(increment)
	gasFeeCap := new(big.Int).Add(increment, priorityFee) /* .Add(increment, big.NewInt(0)) */

	fmt.Printf("%smax fee per gas:%s %s\n", gastParam.ColorGreen, gastParam.ColorReset, gasFeeCap)

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
