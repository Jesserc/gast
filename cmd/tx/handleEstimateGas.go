package transaction

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/internal/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lmittmann/w3"
	w3eth "github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
)

// TryEstimateGas tries to estimate the gas needed to execute a specific transaction based on the current pending state of the backend blockchain. There is no guarantee that this is the true gas limit requiremen
func TryEstimateGas(rpcUrl, from, to, data string, value uint64) (uint64, error) {
	client, err := w3.Dial(rpcUrl)
	if err != nil {
		return 0, fmt.Errorf("failed to dial RPC client: %s", err)
	}
	defer client.Close()

	var (
		fromAddr  = common.HexToAddress(from)
		toAddr    = common.HexToAddress(to)
		amount    = new(big.Int).SetUint64(value)
		bytesData []byte
	)

	var hexData string
	if data != "" {
		// Convert data to hex if it is not
		if !hex.WithOrWithout0xPrefix(data) {
			hexData = hexutil.Encode([]byte(data))
		} else if strings.HasPrefix(data, "0x") {
			hexData = data
		} else {
			hexData = "0x" + data
		}

		bytesData, err = hexutil.Decode(hexData)
		if err != nil {
			return 0, fmt.Errorf("failed to decode data: %s", err)
		}
	}

	var gas uint64
	msg := w3types.Message{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0,
		Value: amount,
		Input: bytesData,
	}

	if err := client.CallCtx(
		context.Background(),
		w3eth.EstimateGas(&msg, nil).Returns(&gas),
	); err != nil {
		return 0, fmt.Errorf("failed estimate gas: %s", err)
	}

	return gas, nil
}
