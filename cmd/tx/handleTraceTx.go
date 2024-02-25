package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
)

type Trace struct {
	Type     string   `json:"type"`
	Depth    int      `json:"depth"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Value    string   `json:"value,omitempty"`
	Input    string   `json:"input"`
	Parent   *Trace   `json:"-"` // Exclude from JSON marshaling
	Children []*Trace `json:"-"`
}

func handleTraceTx(hash, rpcUrl string) (string, error) {

	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		fmt.Println("line 15:", err)
		return "", err
	}

	// params := map[string]interface{}{
	// 	"tracer": "callTracer",
	// 	"tracerConfig": map[string]interface{}{
	// 		"onlyTopCall": false,
	// 		"withLog":     true,
	// 	},
	// }

	var result json.RawMessage
	// err = client.CallContext(context.Background(), &result, "trace_transaction", hash, params)
	err = client.CallContext(context.Background(), &result, "ots_traceTransaction", hash)
	if err != nil {
		return "", err
	}

	var traces []Trace
	if err := json.Unmarshal(result, &traces); err != nil {
		return "", err
	}

	// Organize traces into a hierarchical structure
	traceRoot := buildTraceHierarchy(traces)

	// Pretty print the trace
	printTrace(traceRoot, 0)

	return "output", nil
}

func buildTraceHierarchy(traces []Trace) *Trace {
	var root *Trace
	lastAtDepth := make(map[int]*Trace)

	for i, trace := range traces {
		current := &traces[i] // Get a reference to the trace in the slice

		if trace.Depth == 0 {
			root = current // This is the root trace
		} else {
			parent := lastAtDepth[trace.Depth-1] // Parent is the last trace at the previous depth
			current.Parent = parent
			parent.Children = append(parent.Children, current)
		}

		lastAtDepth[trace.Depth] = current // Update the last trace at this depth
	}

	return root // Return the root of the trace hierarchy
}

func printTrace(trace *Trace, indentLevel int) {
	indent := strings.Repeat(" ", indentLevel*4)
	fmt.Printf("%sType: %s, From: %s, To: %s, Depth: %d\n", indent, trace.Type, trace.From, trace.To, trace.Depth)

	for _, child := range trace.Children {
		printTrace(child, indentLevel+1) // Recursively print children
	}
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
/*
func prettyPrintTraces(traces []TraceDetails, indentLevel int) {
	if traces == nil {
		return
	}
	var output string
	for i, trace := range traces {
		output += fmt.Sprintf("%sTrace %d:\n", strings.Repeat("\t", indentLevel), i+1)
		output += fmt.Sprintf("%s\tCall Type: %s\n", strings.Repeat("\t", indentLevel), trace.Action.CallType)
		output += fmt.Sprintf("%s\tFrom: %s\n", strings.Repeat("\t", indentLevel), trace.Action.From)
		output += fmt.Sprintf("%s\tTo: %s\n", strings.Repeat("\t", indentLevel), trace.Action.To)
		output += fmt.Sprintf("%s\tGas: %s\n", strings.Repeat("\t", indentLevel), trace.Action.Gas)
		output += fmt.Sprintf("%s\tValue: %s\n", strings.Repeat("\t", indentLevel), trace.Action.Value)
		output += fmt.Sprintf("%s\tInput: %s\n", strings.Repeat("\t", indentLevel), trace.Action.Input)
		output += fmt.Sprintf("%s\tOutput: %s\n", strings.Repeat("\t", indentLevel), trace.Result.Output)
		output += fmt.Sprintf("%s\tGas Used: %s\n", strings.Repeat("\t", indentLevel), trace.Result.GasUsed)

	}
	fmt.Println(output)
}
*/

/*
func run(rpcURL, transactionHash string)

	// JSON-RPC request payload
	requestData := fmt.Sprintf(`{"jsonrpc":"2.0","method":"trace_transaction","params":["%s", {}],"id":1}`, transactionHash)

	// Send HTTP POST request to the RPC endpoint
	response, err := http.Post(rpcURL, "application/json", strings.NewReader(requestData))
	if err != nil {
		fmt.Println("Error sending RPC request:", err)
		return
	}
	defer response.Body.Close()

	// Parse the JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Check for RPC errors
	if errMsg, ok := result["error"].(map[string]interface{}); ok {
		code := errMsg["code"].(float64)
		message := errMsg["message"].(string)
		fmt.Println("RPC error:", code, message)
		return
	}

	// Extract the trace result from the response
	traceResult := result["result"].(map[string]interface{})

	// Print the trace result or handle it as needed
	fmt.Println("Trace result:", traceResult)
}*/
