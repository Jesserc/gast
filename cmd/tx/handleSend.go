package transaction

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/Jesserc/gast/internal/hex"
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

// SendTransaction sends an Ethereum transaction.
func SendTransaction(rpcURL, to, data, privateKey string, gasLimit, wei uint64) (string, error) {
	// Connect to the Ethereum client
	client, err := w3.Dial(rpcURL)
	if err != nil {
		return "", fmt.Errorf("failed to dial RPC client: %s", err)
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
			return "", fmt.Errorf("failed to get chain ID: %s", err)
		} else if errs[1] != nil {
			return "", fmt.Errorf("failed to get base fee: %s", err)
		} else if errs[2] != nil {
			return "", fmt.Errorf("failed get priority fee: %s", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("failed RPC request: %s", err)
	}

	log.Info("base fee", "value", baseFee)
	log.Info("priority fee", "value", priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(&baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, &priorityFee) /* .Add(increment, big.NewInt(0)) */
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

	var pendingNonce string
	if err := client.CallCtx(
		context.Background(),
		rpcfactory.PendingNonceAt(fromAddress).Returns(&pendingNonce),
	); err != nil {
		return "", fmt.Errorf("failed to get nonce: %s", err)
	}

	nonce, err := hexutil.DecodeUint64(pendingNonce)
	if err != nil {
		return "", fmt.Errorf("failed to parse nonce: %s", err)
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
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonce,
		GasTipCap: &priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     amount,
		Data:      bytesData,
	}), nil
	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %s", err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(int64(chainID))), ecdsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %s", err)
	}

	fmt.Println() // spacing
	log.Info("Transaction sent")

	var hash common.Hash
	if err := client.CallCtx(
		context.Background(),
		w3eth.SendTx(signedTx).Returns(&hash),
	); err != nil {
		return "", fmt.Errorf("failed to send transaction: %s", err)
	}

	var transactionURL string
	for id, explorer := range gastParams.NetworkExplorers {
		if chainID == id {
			transactionURL = fmt.Sprintf("%vtx/%v", explorer, signedTx.Hash().Hex())
			break
		}
	}
	return transactionURL, nil
}
