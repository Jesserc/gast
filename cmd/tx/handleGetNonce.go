package transaction

import (
	"context"
	"fmt"

	rpcfactory "github.com/Jesserc/gast/internal/rpc_factory"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lmittmann/w3"
)

func GetNonce(address, rpcUrl string) (uint64, uint64, error) {
	client, err := w3.Dial(rpcUrl)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to dial RPC client: %s", err)
	}
	defer client.Close()

	var nonce string
	if err := client.CallCtx(
		context.Background(),
		rpcfactory.PendingNonceAt(w3.A(address)).Returns(&nonce),
	); err != nil {
		return 0, 0, fmt.Errorf("failed to get next nonce: %s", err)
	}

	nextNonce, err := hexutil.DecodeUint64(nonce)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid value received: %s", err)
	}

	var currentNonce uint64
	if nextNonce > 0 {
		currentNonce = nextNonce - 1
	}

	return currentNonce, nextNonce, nil
}
