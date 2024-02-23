package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func handleTraceTx(hash, rpcUrl string) (any, error) {
	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		fmt.Println("line 15:", err)
		return "", err
	}

	params := map[string]any{
		// "tracer":           "prestateTracer",
		"onlyTopCall":      true,
		"disableStack":     true,
		"disableStorage":   true,
		"disableMemory":    true,
		"enableReturnData": false,
		"diffMode":         false,
		"tracer":           "callTracer",

		// "enableMemory":     "true",
		// "disableStack":     "false",
		// "disableStorage":   "false",
		// "enableReturnData": "true",
	}

	// t := tracers.TraceConfig{
	// 	Config:       nil,
	// 	Tracer:       nil,
	// 	Timeout:      nil,
	// 	Reexec:       nil,
	// 	TracerConfig: nil,
	// }

	var response json.RawMessage
	err = client.Call(&response, "debug_traceTransaction", hash, params, map[string]any{"tracer": "prestateTracer" /* "disableStack": true, "disableStorage": true*/})

	if err != nil {
		fmt.Println("line 21:", err)
		return "", err
	}

	v, err := json.MarshalIndent(response, " ", " ")
	if err != nil {
		return nil, err
	}

	return string(v), nil
}

// func handleTraceTx(hash, rpcUrl string) (any, error) {
// 	client, err := rpc.Dial(rpcUrl)
// 	if err != nil {
// 		fmt.Println("line 15:", err)
// 		return "", err
// 	}
//
// 	t := "callTracer"
// 	lc := logger.Config{
// 		EnableMemory:     false,
// 		DisableStack:     true,
// 		DisableStorage:   true,
// 		EnableReturnData: false,
// 		Debug:            false,
// 		Limit:            0,
// 		Overrides:        nil,
// 	}
// 	tc := tracers.TraceConfig{
// 		Config:       &lc,
// 		Tracer:       &t,
// 		Timeout:      nil,
// 		Reexec:       nil,
// 		TracerConfig: nil,
// 	}
//
// 	message, err := TraceTransaction(context.Background(), client, common.HexToHash(hash), &tc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	d, err := json.MarshalIndent(message, " ", "\t")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return string(d), err
// }
//
// func TraceTransaction(ctx context.Context, ec *rpc.Client, txhash common.Hash, config *tracers.TraceConfig) (json.RawMessage, error) {
// 	var result json.RawMessage
// 	err := ec.CallContext(ctx, &result, "debug_traceTransaction", txhash, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
