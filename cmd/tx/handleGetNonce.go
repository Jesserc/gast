package transaction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func GetNonce(address, rpcUrl string) (uint64, uint64) {
	client, err := ethclient.Dial(rpcUrl)
	defer client.Close()
	if err != nil {
		log.Crit("Failed to dial RPC client", "error", err)
	}

	nextNonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		log.Crit("Failed to get nonce", "error", err)
	}

	var currentNonce uint64
	if nextNonce > 0 {
		currentNonce = nextNonce - 1
	}

	return currentNonce, nextNonce
}
