package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jesserc/gast/internal/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
	"github.com/mitchellh/go-homedir"
)

func SendBlobTX(rpcURL, toAddress, data, privateKey, saveBlobDir string) (string, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", fmt.Errorf("failed to dial RPC client: %s", err)
	}
	defer client.Close()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %s", err)
	}

	var Blob [131072]byte

	// Convert data to hex format
	var bytesData []byte
	if data != "" {
		if hex.WithOrWithout0xPrefix(data) {
			if !strings.HasPrefix(data, "0x") {
				data = "0x" + data // add `0x` prefix if it doesn't have
			}
			bytesData, err = hexutil.Decode(data)
			if err != nil {
				return "", fmt.Errorf("failed to decode data: %s", err)
			}

			copy(Blob[:], bytesData)

		} else {
			copy(Blob[:], data) // if it's not hex, just copy into blob (it'll be converted to bytes by the copy fn)
		}

	}

	BlobCommitment, err := kzg4844.BlobToCommitment(Blob)
	if err != nil {
		return "", fmt.Errorf("failed to compute blob commitment: %s", err)
	}

	BlobProof, err := kzg4844.ComputeBlobProof(Blob, BlobCommitment)
	if err != nil {
		return "", fmt.Errorf("failed to compute blob proof: %s", err)
	}

	sidecar := types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{Blob},
		Commitments: []kzg4844.Commitment{BlobCommitment},
		Proofs:      []kzg4844.Proof{BlobProof},
	}

	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %s", err)
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key to ECDSA: %s", err)
	}

	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %s", err)
	}

	tx, err := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(chainID),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(1e10),
		GasFeeCap:  uint256.NewInt(20e10),
		Gas:        250000,
		To:         common.HexToAddress(toAddress),
		Value:      uint256.NewInt(0),
		Data:       nil,
		BlobFeeCap: uint256.NewInt(3e10), // 30 gwei
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    &sidecar,
	}), nil
	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %s", err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %s", err)
	}

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", fmt.Errorf("failed to send transaction: %s", err)
	}

	if saveBlobDir != "" {
		d, err := json.MarshalIndent(signedTx, "", "\t")
		if err != nil {
			log.Error("Failed to marshal indent transaction", "error", err)
		}

		hd, err := homedir.Dir()
		if err != nil {
			log.Error("Failed to get home directory", "error", err)
		}

		path := filepath.Join(hd, saveBlobDir)
		if err = os.MkdirAll(path, 0755); err != nil {
			log.Error("Failed to create directory", "error", err)
		}

		n := fmt.Sprintf("blob_%v", signedTx.Hash().Hex()[0:6])
		f, err := os.Create(filepath.Join(path, n))
		if err != nil {
			log.Error("Failed to create file", "error", err)
		}

		_, err = f.Write(d)
		if err != nil {
			log.Error("Failed to write blob tx details to file", "error", err)
		} else {
			log.Info("Blob transaction details saved", "filepath", f.Name())
		}
	}
	txHash := signedTx.Hash().Hex()

	return txHash, nil
}
