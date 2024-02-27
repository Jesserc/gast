package transaction

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// handleVerifySig verifies the signature against the provided public key bytes and hash.
func handleVerifySig(signature, address, hashStr string) bool {
	sig := []byte(signature)

	// Adjust signature to standard format (remove Ethereum's recovery ID)
	sig[64] = sig[64] - 27

	// Convert hash string to common.Hash (github.com/ethereum/go-ethereum/common)
	hash := common.HexToHash(hashStr)

	// Recover the public key bytes from the signature
	sigPublicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}
	ecdsaPublicKey, err := crypto.UnmarshalPubkey(sigPublicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Derive the address from the recovered public key
	rAddress := crypto.PubkeyToAddress(*ecdsaPublicKey)

	// Check if the recovered address matches the provided address
	equals := strings.EqualFold(rAddress.String(), address)

	// Verify if the recovered public key matches the provided public key bytes
	return equals
}
