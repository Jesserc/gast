package transaction

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func handleSignMessageHash(rpcUrl, to, data, privateKey string, wei, nonce uint64) (string, error) {
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

	k, err := hexutil.Decode(privateKey)
	if err != nil {
		return "", err
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(k)
	if err != nil {
		return "", err
	}

	h := "0x" + hex.EncodeToString([]byte(data))

	bytesData, err := hexutil.Decode(h)
	if err != nil {
		return "", err
	}

	txData := types.LegacyTx{
		Nonce: nonce,
		To:    &toAddr,
		Value: amount,
		Data:  bytesData,
	}

	tx := types.NewTx(&txData)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	hash := hexutil.Encode(signedTx.Hash().Bytes())

	fmt.Println("hash:", hash)
	return hash, nil
}

func handleSignMessageRaw(rpcUrl, to, data, privateKey string, wei, nonce uint64) (string, error) {
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

	k, err := hexutil.Decode(privateKey)
	if err != nil {
		return "", err
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(k)
	if err != nil {
		return "", err
	}

	h := "0x" + hex.EncodeToString([]byte(data))

	bytesData, err := hexutil.Decode(h)
	if err != nil {
		return "", err
	}

	txData := types.LegacyTx{
		Nonce: nonce,
		To:    &toAddr,
		Value: amount,
		Data:  bytesData,
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

	fmt.Println("rlp encoded signed message:", rawTxRLPHex)

	return rawTxRLPHex, nil
}

/*
	var buf bytes.Buffer
	// var buf2 bytes.Buffer
	signedTx.EncodeRLP(&buf)
	fmt.Println("buf 2", buf)

	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	fmt.Println("rlp encoded signed message:\n", rawTxRLPHex)

	signedTx.DecodeRLP(rlp.NewStream(&buf, 0))
	fmt.Println("buf 2", buf)

0xd7826eefd8fca691249af9775d18626ed87c3892fab580c68332905c0d226f771e6e60713f141e431ba76b26467bd1ce651cb2c2d32d7e0225483ac67f24e6b01b
*/
