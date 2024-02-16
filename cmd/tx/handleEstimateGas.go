package transaction

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

var (
	ctx = context.Background()
)

func HandleEstimateGas(rpcUrl, from, to, data string, wei uint64) (uint64, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return 0, err
	}

	var (
		fromAddr = common.HexToAddress(from)
		toAddr   = common.HexToAddress(to)
		value    = new(big.Int).SetUint64(wei)
	)

	bytesData, err := hexutil.Decode(data)
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: value,
		Data:  bytesData,
	}

	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}

	return gas, nil
}

/*
			type CallMsg struct {
				From      common.Address  // the sender of the 'transaction'
				To        *common.Address // the destination contract (nil for contract creation)
				Gas       uint64          // if 0, the call executes with near-infinite gas
				GasPrice  *big.Int        // wei <-> gas exchange ratio
				GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
				GasTipCap *big.Int        // EIP-1559 tip per gas.
				Value     *big.Int        // amount of wei sent along with the call
				Data      []byte          // input data, usually an ABI-encoded contract method invocation

				AccessList types.AccessList // EIP-2930 access list.
		var (
				fromAddr = common.HexToAddress("0x8D97689C9818892B700e27F316cc3E41e17fBeb9")
				toAddr   = common.HexToAddress("0xd3CdA913deB6f67967B99D67aCDFa1712C293601")
		)

	// data := "0xde725e890000000000000000000000000f93ae9f3b81c12cbc009e8f0d4a4f4f044df3040000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d0000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000"
	// 0xde725e890000000000000000000000000f93ae9f3b81c12cbc009e8f0d4a4f4f044df3040000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d0000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000
	// 0x12514bba0000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000
	// 5524939


	--url https://nd-422-757-666.p2pify.com/0a9d79d93fb2f4a4b1e04695da2b77a7/ \
	--header 'accept: application/json' \
	--header 'content-type: application/json' \
	--data '
	{
	"id": 1,
	"jsonrpc": "2.0",
	"method": "eth_estimateGas",
	"params": [
	{
	"from": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
	"to": "0xbe0eb53f46cd790cd13851d5eff43d12404d33e8", "data":"0xde725e890000000000000000000000000f93ae9f3b81c12cbc009e8f0d4a4f4f044df3040000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d0000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000"
	},
	"latest"
	]
	}'
	{"jsonrpc":"2.0","id":1,"result":"0x5630"}

*/
