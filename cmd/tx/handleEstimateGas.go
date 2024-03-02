package transaction

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TryEstimateGas tries to estimate the gas needed to execute a specific transaction based on the current pending state of the backend blockchain. There is no guarantee that this is the true gas limit requiremen
func TryEstimateGas(rpcUrl, from, to, data string, value uint64) (uint64, error) {
	client, err := ethclient.Dial(rpcUrl)
	defer client.Close()
	if err != nil {
		return 0, err
	}

	var ctx = context.Background()

	var (
		fromAddr  = common.HexToAddress(from)
		toAddr    = common.HexToAddress(to)
		amount    = new(big.Int).SetUint64(value)
		bytesData []byte
	)

	if data != "" {
		if ok := strings.HasPrefix(data, "0x"); !ok {
			data = hexutil.Encode([]byte(data))
		}

		bytesData, err = hexutil.Decode(data)
		if err != nil {
			return 0, err
		}
	}

	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: amount,
		Data:  bytesData,
	}

	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}

	return gas, nil
}
