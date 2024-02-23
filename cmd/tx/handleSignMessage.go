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
