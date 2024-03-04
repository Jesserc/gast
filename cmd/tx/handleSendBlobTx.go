package transaction

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

func SendBlobTX(rpcURL, data, privateKey, toAddress string) string {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Error("Failed to dial RPC client", "error", err)
		return ""
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Error("Failed to get chain ID", "error", err)
		return ""
	}

	var Blob [131072]byte
	if strings.HasPrefix(data, "0x") {
		bytesData, err := hexutil.Decode(data)
		if err != nil {
			log.Error("Failed to decode data", "error", err)
			return ""
		}
		copy(Blob[:], bytesData)
	} else {
		copy(Blob[:], data)
	}

	BlobCommitment, err := kzg4844.BlobToCommitment(Blob)
	if err != nil {
		log.Error("Failed to compute blob commitment", "error", err)
		return ""
	}

	BlobProof, err := kzg4844.ComputeBlobProof(Blob, BlobCommitment)
	if err != nil {
		log.Error("Failed to compute blob proof", "error", err)
		return ""
	}

	sidecar := types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{Blob},
		Commitments: []kzg4844.Commitment{BlobCommitment},
		Proofs:      []kzg4844.Proof{BlobProof},
	}

	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		log.Error("Failed to decode private key", "error", err)
		return ""
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		log.Error("Failed to convert private key to ECDSA", "error", err)
		return ""
	}

	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error("Failed to get nonce", "error", err)
		return ""
	}

	tx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(chainID),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(1e10),
		GasFeeCap:  uint256.NewInt(2e10),
		Gas:        250000,
		To:         common.HexToAddress(toAddress),
		Value:      uint256.NewInt(0),
		Data:       nil,
		BlobFeeCap: uint256.NewInt(3e10), // 30 gwei
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    &sidecar,
	})

	if err != nil {
		log.Error("Failed to create transaction", "error", err)
		return ""
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		log.Error("Failed to sign transaction", "error", err)
		return ""
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send transaction", "error", err)
		return ""
	}

	d, err := json.MarshalIndent(signedTx, "", "\t")
	if err != nil {
		log.Error("Failed to marhshal indent transaction", "error", err)
	}

	// homeDir, err := homedir.Dir()
	// if err != nil {
	// 	return ""
	// }
	//

	f, err := os.Create("blob")
	if err != nil {
		log.Error("Failed to create blob file", "error", err)
	}
	defer f.Close()

	_, err = f.Write(d)
	if err != nil {
		log.Error("Failed write to blob file", "error", err)
	}

	log.Info("Blob transaction details saved", "file name", f.Name())

	txHash := signedTx.Hash().Hex()

	return txHash
}

/*package transaction

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

func SendBlobTX(rpcURL, data, privateKey, toAddress string) (string, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Error("Failed to dial Ethereum client", "rpcURL", rpcURL, "error", err)
		return "", err // Return here as we can't proceed without a client
	}

	// Blob
	var Blob [131072]byte
	if strings.HasPrefix(data, "0x") {
		bytesData, err := hexutil.Decode(data)
		if err != nil {
			return "", err
		}

		copy(Blob[:], bytesData[:]) // copy data into blob
	} else {
		copy(Blob[:], data[:]) // copy data into blob
	}

	BlobCommitment, err := kzg4844.BlobToCommitment(Blob)
	if err != nil {
		return "", err
	}

	BlobProof, err := kzg4844.ComputeBlobProof(Blob, BlobCommitment)
	if err != nil {
		return "", err
	}
	sidecar := types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{Blob},
		Commitments: []kzg4844.Commitment{BlobCommitment},
		Proofs:      []kzg4844.Proof{BlobProof},
	}

	// Decode private key
	pKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return "", err
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		return "", err
	}

	publicKey := ecdsaPrivateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA pubKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Crit("Failed to get nonce", "err", err)
	}

	tx, err := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(chainID),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(1e10),
		GasFeeCap:  uint256.NewInt(2e10),
		Gas:        250000,
		To:         common.HexToAddress(toAddress),
		Value:      uint256.NewInt(0),
		Data:       nil,
		BlobFeeCap: uint256.NewInt(3e10), // 30 gwei
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    &sidecar,
	}), nil
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewCancunSigner(tx.ChainId()), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	d, err := json.MarshalIndent(signedTx, "", "\t")
	if err != nil {
		return "", err
	}

	f, err := os.Create("blob")
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(d)
	if err != nil {
		return "", err
	}

	log.Info("Blob transaction details saved", "file name", f.Name())
	return signedTx.Hash().Hex(), nil
}
*/
