package transaction

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func handleSignMessage(rpcUrl, to, data, privateKey string, wei, nonce uint64) (string, error) {
	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(wei)

	k, err := hexutil.Decode(privateKey)
	if err != nil {
		// fmt.Println(err)

		return "", err
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(k)
	if err != nil {
		// fmt.Println(err)

		return "", err
	}

	h := "0x" + hex.EncodeToString([]byte(data))

	bytesData, err := hexutil.Decode(h)
	if err != nil {
		fmt.Println(err)

		return "", err
	}

	txData := types.LegacyTx{
		Nonce: nonce,
		To:    &toAddr,
		Value: amount,
		Data:  bytesData,
	}

	tx := types.NewTx(&txData)
	fmt.Println("transaction:", hexutil.Encode(tx.Data()))

	id := big.NewInt(11155111)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(id), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	// var buf bytes.Buffer
	// // var buf2 bytes.Buffer
	// signedTx.EncodeRLP(&buf)
	// fmt.Println("buf 2", buf)
	//
	// rawTxRLPHex := hex.EncodeToString(buf.Bytes())
	//
	// fmt.Println("rlp encoded signed message:\n", rawTxRLPHex)
	// signedTx.DecodeRLP(rlp.NewStream(&buf, 0))
	// fmt.Println("buf 2", buf)
	fmt.Println(signedTx.Hash())
	return "", nil
}

/*
0xd7826eefd8fca691249af9775d18626ed87c3892fab580c68332905c0d226f771e6e60713f141e431ba76b26467bd1ce651cb2c2d32d7e0225483ac67f24e6b01b
*/
