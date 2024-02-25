package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
)

type TopCall struct {
	From    string
	Gas     string
	GasUsed string
	Input   string
	To      string
	Type    string
	Value   string
}

func handleTraceTx(hash, rpcUrl string) (string, error) {

	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		fmt.Println("line 15:", err)
		return "", err
	}

	params := map[string]interface{}{
		"tracer": "callTracer",
		"tracerConfig": map[string]interface{}{
			"onlyTopCall": false,
			"withLog":     true,
		},
	}

	var result json.RawMessage
	err = client.CallContext(context.Background(), &result, "trace_transaction", hash, params)
	if err != nil {
		return "", err
	}

	var traces []TraceDetails
	err = json.Unmarshal(result, &traces)
	if err != nil {
		return "", err
	}

	prettyPrintTraces(traces, 0)

	return "output", nil
}

func prettyPrintTraces(traces []TraceDetails, indentLevel int) {
	if traces == nil {
		return
	}
	var output string
	for i, trace := range traces {
		output += fmt.Sprintf("%sSubcall %d:\n", strings.Repeat("  ", indentLevel+i), i+1)
		output += fmt.Sprintf("%sCall Type: %s\n", strings.Repeat("  ", indentLevel+i), trace.Action.CallType)
		output += fmt.Sprintf("%sFrom: %s\n", strings.Repeat("  ", indentLevel+i), trace.Action.From)
		output += fmt.Sprintf("%sTo: %s\n", strings.Repeat("  ", indentLevel+i), trace.Action.To)
		output += fmt.Sprintf("%sGas: %s\n", strings.Repeat("  ", indentLevel+i), trace.Action.Gas)
		output += fmt.Sprintf("%sValue: %s\n", strings.Repeat("  ", indentLevel+i), trace.Action.Value)
		// output += fmt.Sprintf("%s  Input: %s\n", strings.Repeat("  ", indentLevel+i), trace.Action.Input)
		output += fmt.Sprintf("%sOutput: %s\n", strings.Repeat("  ", indentLevel+i), trace.Result.Output)
		output += fmt.Sprintf("%sGas Used: %s\n", strings.Repeat("  ", indentLevel+i), trace.Result.GasUsed)

	}
	fmt.Println(output)
}

// language=json
type TraceDetails struct {
	Action struct {
		CallType string `json:"callType"`
		From     string `json:"from"`
		Gas      string `json:"gas"`
		Input    string `json:"input"`
		To       string `json:"to"`
		Value    string `json:"value"`
	} `json:"action"`
	BlockHash   string `json:"blockHash"`
	BlockNumber int    `json:"blockNumber"`
	Result      struct {
		GasUsed string `json:"gasUsed"`
		Output  string `json:"output"`
	} `json:"result"`
	Subtraces           int    `json:"subtraces"`
	TraceAddress        []int  `json:"traceAddress"`
	TransactionHash     string `json:"transactionHash"`
	TransactionPosition int    `json:"transactionPosition"`
	Type                string `json:"type"`
}

/*func run(rpcURL, transactionHash string) {

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
