package transaction

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/internal/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TryEstimateGas tries to estimate the gas needed to execute a specific transaction based on the current pending state of the backend blockchain. There is no guarantee that this is the true gas limit requiremen
func TryEstimateGas(rpcUrl, from, to, data string, value uint64) (uint64, error) {
	client, err := ethclient.Dial(rpcUrl)
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

	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: amount,
		Data:  bytesData,
	}

	gas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, fmt.Errorf("failed estimate gas: %s", err)
	}

	return gas, nil
}
