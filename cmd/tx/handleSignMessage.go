package transaction

import (
	"encoding/json"
	"fmt"

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

// SignETHMessage signs a message using the provided private key.
func SignETHMessage(message, privateKey string) (string, error) {
	// Convert the private key from hex to ECDSA format
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key to ECDSA: %s", err)
	}

	// Construct the message prefix
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))
	data := []byte(message)

	// Hash the prefix and data using Keccak-256
	hash := crypto.Keccak256Hash(prefix, data)

	// Sign the hashed message
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %s", err)
	}

	// Adjust signature to Ethereum's format
	sig[64] += 27

	// Derive the public key from the private key
	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)

	// Construct the signature response
	res := SignatureResponse{
		Address: fromAddress.String(),
		Msg:     message,
		Sig:     hexutil.Encode(sig),
		Version: "2",
	}

	// Marshal the response to JSON with proper formatting
	resBytes, err := json.MarshalIndent(res, " ", "\t")
	if err != nil {
		return "", fmt.Errorf("failed marshal indent signature response: %s", err)
	}

	return string(resBytes), nil
}
