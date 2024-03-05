package transaction

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNonce(address, rpcUrl string) (uint64, uint64, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to dial RPC client: %s", err)
	}
	defer client.Close()

	nextNonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get nonce: %s", err)
	}

	var currentNonce uint64
	if nextNonce > 0 {
		currentNonce = nextNonce - 1
	}

	return currentNonce, nextNonce, nil
}
