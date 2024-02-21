package transaction

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func handleCreateRawTransaction(rpcUrl, to, data, privateKey string, gasPrice, gasLimit, wei, nonce uint64) (string, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return "", err
	}

	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(wei)
	gPrice := new(big.Int).SetUint64(gasPrice)

	pKey, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return "", err
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(pKey)
	if err != nil {
		return "", err
	}

	hexData := "0x" + hex.EncodeToString([]byte(data))

	bytesData, err := hexutil.Decode(hexData)
	if err != nil {
		return "", err
	}

	txData := types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddr,
		Value:    amount,
		Data:     bytesData,
		GasPrice: gPrice,
		Gas:      gasLimit,
	}

	tx := types.NewTx(&txData)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = signedTx.EncodeRLP(&buf)
	if err != nil {
		return "", err
	}

	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	return rawTxRLPHex, nil
}
