package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignatureResponse represents the structure of the signature response.
type SignatureResponse struct {
	Address string `json:"address,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Sig     string `json:"sig,omitempty"`
	Version string `json:"version,omitempty"`
}

// signMessage signs a message using the provided private key.
func signMessage(message, privateKey string) (string, error) {
	// Convert the private key from hex to ECDSA format
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	// Construct the message prefix
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))
	data := []byte(message)

	// Hash the prefix and data using Keccak-256
	hash := crypto.Keccak256Hash(prefix, data)

	// Sign the hashed message
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	// Adjust signature to Ethereum's format
	sig[64] = sig[64] + 27

	// Derive the public key from the private key
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	rAddress := crypto.PubkeyToAddress(*pub)

	// Construct the signature response
	res := SignatureResponse{
		Address: rAddress.String(),
		Msg:     message,
		Sig:     hexutil.Encode(sig),
		Version: "2",
	}

	// Marshal the response to JSON with proper formatting
	var w bytes.Buffer
	var v bytes.Buffer
	json.NewEncoder(&w).Encode(res)
	json.Indent(&v, w.Bytes(), " ", "\t")

	return v.String(), nil
}

// verifySig verifies the signature against the provided public key bytes and hash.
func verifySig(sig, pubKeyBytes []byte, hash common.Hash) bool {
	// Adjust signature to standard format
	sig[64] = sig[64] - 27

	// Recover the public key from the signature
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}
	pub, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the recovered public key matches the provided public key bytes
	fmt.Println(bytes.Equal(sigPublicKey, pubKeyBytes))

	// Derive the address from the recovered public key
	rAddress := crypto.PubkeyToAddress(*pub)

	// Print the recovered address
	fmt.Println("Recovered address:", rAddress)

	// Verify if the recovered public key matches the provided public key bytes
	return bytes.Equal(sigPublicKey, pubKeyBytes)
}

// func handleSignMessageHash(data, privateKey string) (string, error) {
// 	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	fmt.Println(crypto.PubkeyToAddress(*ecdsaPrivateKey.Public().(*ecdsa.PublicKey))) // 0x571B102323C3b8B8Afb30619Ac1d36d85359fb84 // correct address but verification fails, why?
//
// 	// Construct the message prefix
// 	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
//
// 	// Concatenate the prefix and data
// 	message := append([]byte(prefix), []byte(data)...)
//
// 	// Hash the message using Keccak-256
// 	hash := crypto.Keccak256(message)
//
// 	// Sign the hashed message
// 	signature, err := crypto.Sign(hash, ecdsaPrivateKey)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	// Ensure that the signature verification is successful
// 	match := crypto.VerifySignature(crypto.CompressPubkey(ecdsaPrivateKey.Public().(*ecdsa.PublicKey)), hash, signature)
// 	fmt.Println(match)
//
// 	// Modify the signature (if needed)
// 	signature[64] = 0x1c
//
// 	return hexutil.Encode(signature), nil
// }

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
