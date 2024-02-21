package transaction

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

func estimateGas(rpcUrl, from, to, data string, value uint64) (uint64, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return 0, err
	}

	var ctx = context.Background()

	var (
		fromAddr = common.HexToAddress(from)
		toAddr   = common.HexToAddress(to)
		amount   = new(big.Int).SetUint64(value)
	)

	bytesData, err := hexutil.Decode(data)
	if err != nil {
		return 0, err
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
