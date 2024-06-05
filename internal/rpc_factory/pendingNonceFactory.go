package rpcfactory

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
)

type getPendingNonceFactory struct {
	address common.Address
	returns *string
}

func (f *getPendingNonceFactory) Returns(res *string) w3types.RPCCaller {
	f.returns = res
	return f
}

func (f *getPendingNonceFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_getTransactionCount",
		Args:   []any{f.address, "pending"},
		Result: f.returns,
	}, nil
}

func (f *getPendingNonceFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}

	return nil
}

func PendingNonceAt(address common.Address) w3types.RPCCallerFactory[string] {
	return &getPendingNonceFactory{
		address: address,
	}
}
