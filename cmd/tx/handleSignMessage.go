package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func handleSignMessageHash(data, privateKey string) (string, error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(ecdsaPrivateKey)

		return "", err
	}
	pk := ecdsaPrivateKey.Public().(*ecdsa.PublicKey)

	fmt.Println(crypto.PubkeyToAddress(*pk))
	fmt.Println(len(data))
	//
	// message := "\u0019Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(data)), 16) + data
	// hash := crypto.Keccak256Hash([]byte(message))
	// fmt.Println(hash.Hex())

	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data)))

	// Hash the message using Keccak-256
	hash := crypto.Keccak256(prefix, []byte(data))

	signature, err := crypto.Sign(hash, ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	sigPublicKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(crypto.PubkeyToAddress(*sigPublicKey))
	signature[64] = 0x1c
	return hexutil.Encode(signature), nil
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

// handleSignMessageHash
// client, err := ethclient.Dial(rpcUrl)
	// if err != nil {
	// 	return "", err
	// }
	//
	// ctx := context.Background()
	// chainID, err := client.ChainID(ctx)
	// if err != nil {
	// 	return "", err
	// }
	//
	// toAddr := common.HexToAddress(to)
	// amount := new(big.Int).SetUint64(wei)


	// h := "0x" + hex.EncodeToString([]byte(data))
	//
	// bytesData, err := hexutil.Decode(h)
	// if err != nil {
	// 	return "", err
	// }

	// txData := types.LegacyTx{
	// 	Nonce: nonce,
	// 	To:    &toAddr,
	// 	Value: amount,
	// 	Data:  bytesData,
	// }
	//
	// tx := types.NewTx(&txData)
	//
	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ecdsaPrivateKey)
	// if err != nil {
	// 	return "", err
	// }
	//
	// hash := hexutil.Encode(signedTx.Hash().Bytes())
*/
