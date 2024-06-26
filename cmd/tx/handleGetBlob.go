package transaction

import (
	"context"
	"encoding/json"
	"fmt"

	eth2api "github.com/attestantio/go-eth2-client/api"
	eth2http "github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/rs/zerolog"
)

func GetBlob(rpcUrl, blockRootOrSlotNumber, kzgCommitment string) (string, error) {
	var err error
	ctx, cancel := context.WithCancel(context.Background())

	client, err := eth2http.New(ctx,
		// WithAddress supplies the address of the beacon node, as a URL.
		eth2http.WithAddress(rpcUrl),
		// LogLevel supplies the level of logging to carry out.
		eth2http.WithLogLevel(zerolog.WarnLevel),
	)
	if err != nil {
		cancel()
		return "", fmt.Errorf("failed to create eth-2 http client: %v", err)
	}

	log.Warn("", "Connected to", client.Address())

	var blob any

	httpClient := client.(*eth2http.Service)
	blobSideCarResponse, err := httpClient.BlobSidecars(ctx, &eth2api.BlobSidecarsOpts{
		Common: eth2api.CommonOpts{ /*Timeout: 30 * time.Second*/ },
		Block:  blockRootOrSlotNumber,
	})
	if err != nil {
		cancel()
		return "", fmt.Errorf("failed to establish BlobSidecarsProvider: %v, ensure parameters are correct", err)
	}

	log.Info("Fetching blob...")

	for i, bs := range blobSideCarResponse.Data {
		ksc, err := hexutil.Decode(kzgCommitment)
		if err != nil {
			cancel()
			return "", fmt.Errorf("failed to decode hex kzg commitment: %v", err)
		}

		kc := deneb.KZGCommitment(ksc)
		if bs.KZGCommitment == kc {
			blob = blobSideCarResponse.Data[i] // set blob
			break
		}
	}

	// Cancelling the context passed to New() frees up resources held by the
	// client, closes connections, clears handlers, etc.
	cancel()

	blobBytes, err := json.MarshalIndent(blob, "", "\t")
	if err != nil {
		return "", fmt.Errorf("failed to marshal blob: %v", err)
	}
	return string(blobBytes), nil
}
