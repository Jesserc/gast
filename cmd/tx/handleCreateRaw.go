package transaction

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// CreateRawTransaction creates a raw Ethereum transaction.
func CreateRawTransaction(rpcURL, to, data, privateKey string, gasLimit, wei uint64) string {
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
		log.Crit("Failed to get priority fee", "error", err) // TODO: should this not be critical? i.e, it shouldn't stop execution here...
	}
	log.Info("priority fee", "value", priorityFee)

	// Calculate gas fee cap with 2 Gwei margin
	increment := new(big.Int).Add(baseFee, big.NewInt(2*params.GWei))
	gasFeeCap := new(big.Int).Add(increment, priorityFee) /* .Add(increment, big.NewInt(0)) */
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

	// Convert data to hex format
	var hexData string
	var bytesData []byte
	if data != "" {
		if !utils.IsHexWithOrWithout0xPrefix(data) {
			hexData = hexutil.Encode([]byte(data))
		} else if strings.HasPrefix(data, "0x") {
			hexData = data
		} else {
			hexData = "0x" + data
		}

		bytesData, err = hexutil.Decode(hexData)
		if err != nil {
			log.Crit("Failed to decode data", "error", err)
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
		log.Crit("Failed to create transaction", "error", err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		log.Crit("Failed to sign transaction", "error", err)
	}

	// Encode signed transaction to RLP hex
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)
	if err != nil {
		log.Crit("Failed get rlp encoded raw transaction", "error", err)
	}

	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	return rawTxRLPHex
}
