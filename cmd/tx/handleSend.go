package transaction

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/Jesserc/gast/internal/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// SendTransaction sends an Ethereum transaction.
func SendTransaction(rpcURL, to, data, privateKey string, gasLimit, wei uint64) (string, error) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", fmt.Errorf("failed to dial RPC client: %s", err)
	}
	defer client.Close()

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %s", err)
	}

	// Get base fee
	baseFee, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get base fee: %s", err)
	}
	log.Info("base fee", "value", baseFee)

	// Get priority fee
	priorityFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed get priority fee: %s", err)
	}
	log.Info("priority fee", "value", priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, priorityFee) /* .Add(increment, big.NewInt(0)) */
	log.Info("max fee per gas", "value", gasFeeCap)

	// Decode private key
	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %s", err)
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key to ECDSA: %s", err)
	}

	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %s", err)
	}

	// Convert data to hex format
	var hexData string
	var bytesData []byte
	if data != "" {
		if !hex.WithOrWithout0xPrefix(data) {
			hexData = hexutil.Encode([]byte(data))
		} else if strings.HasPrefix(data, "0x") {
			hexData = data
		} else {
			hexData = "0x" + data
		}

		bytesData, err = hexutil.Decode(hexData)
		if err != nil {
			return "", fmt.Errorf("failed to decode data: %s", err)
		}
	}

	// Create transaction data
	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(wei)

	tx, err := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     amount,
		Data:      bytesData,
	}), nil
	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %s", err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %s", err)
	}

	fmt.Println() // spacing
	log.Info("Transaction sent")

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", fmt.Errorf("failed to send transaction: %s", err)
	}

	var transactionURL string
	for id, explorer := range gastParams.NetworkExplorers {
		if chainID.Uint64() == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, signedTx.Hash().Hex())
			break
		}
	}
	return transactionURL, nil
}
