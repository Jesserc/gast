package transaction

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// handleVerifySig verifies the signature against the provided public key bytes and hash.
func handleVerifySig(sigString, pubKey string, hash common.Hash) bool {
	crypto.DecompressPubkey([]byte(pubKey))

	pubKeyBytes := []byte(pubKey)
	sig := []byte(sigString)

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
