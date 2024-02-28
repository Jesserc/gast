package transaction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func handleGetNonce(address, rpcUrl string) (uint64, uint64, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return 0, 0, err
	}

	nextNonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return 0, 0, err
	}

	var currentNonce uint64
	if nextNonce > 0 {
		currentNonce = nextNonce - 1
	}
	
	return currentNonce, nextNonce, nil
}
