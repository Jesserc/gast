package rpcfactory

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
)

type getTraceTransactionFactory struct {
	hash    string
	returns *json.RawMessage
}

func (f *getTraceTransactionFactory) Returns(res *json.RawMessage) w3types.RPCCaller {
	f.returns = res
	return f
}

func (f *getTraceTransactionFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "ots_traceTransaction",
		Args:   []any{f.hash},
		Result: f.returns,
	}, nil
}

func (f *getTraceTransactionFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}

	return nil
}

func OtsTraceTransaction(hash string) w3types.RPCCallerFactory[json.RawMessage] {
	return &getTraceTransactionFactory{
		hash: hash,
	}
}
