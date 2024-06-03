package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/Jesserc/gast/cmd/gastParams"
	rpcfactory "github.com/Jesserc/gast/internal/rpc_factory"
	"github.com/ethereum/go-ethereum/log"
	"github.com/lmittmann/w3"
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

func TraceTx(hash, rpcUrl string) (*Trace, error) {
	var (
		client *w3.Client
		err    error
	)

	if rpcUrl == "" {
		client, err = w3.Dial("https://rpc.builder0x69.io/")
		if err != nil {
			return nil, fmt.Errorf("failed to dial RPC client: %s", err)
		}
		defer client.Close()
	} else {
		client, err = w3.Dial(rpcUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to dial RPC client: %s", err)
		}
		defer client.Close()
	}

	var result json.RawMessage
	if err := client.CallCtx(
		context.Background(),
		rpcfactory.OtsTraceTransaction(hash).Returns(&result),
	); err != nil {
		return nil, fmt.Errorf("failed to trace transaction: %s", err)
	}

	var traces []Trace
	if err = json.Unmarshal(result, &traces); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trace result to Go type []Trace: %s", err)
	}

	rootTrace := buildTraceHierarchy(traces)
	return rootTrace, nil
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

	return root
}

func printTrace(trace *Trace, indentLevel int, isLastChild bool, prefix string) {
	var indent, currentPrefix string
	if indentLevel > 0 {
		indent = strings.Repeat("", indentLevel-1) // Basic indentation for hierarchy level
		if isLastChild {
			currentPrefix = prefix + "└── " + gastParams.ColorGreen + "← " + gastParams.ColorReset
			prefix += "    " // Extend the prefix for child traces without a connecting line
		} else {
			currentPrefix = prefix + "├── " + gastParams.ColorGreen + "← " + gastParams.ColorReset
			prefix += "│   " // Add a vertical line for child traces
		}
	}
	formattedInput := formatInput(trace.Input) // Format the input field
	fmt.Printf("%s%s%sType:%s %s, %sFrom:%s %s, %sTo:%s %s, %sDepth:%s %d, %sValue:%s %s, %sInput:%s [%s]\n",
		indent, currentPrefix,
		gastParams.ColorGreen, gastParams.ColorReset,
		trace.Type,
		gastParams.ColorGreen, gastParams.ColorReset,
		trace.From,
		gastParams.ColorGreen, gastParams.ColorReset,
		trace.To,
		gastParams.ColorGreen, gastParams.ColorReset,
		trace.Depth,
		gastParams.ColorGreen, gastParams.ColorReset,
		hexToEther(trace.Value),
		gastParams.ColorGreen, gastParams.ColorReset,
		formattedInput,
	)

	for i, child := range trace.Children {
		printTrace(child, indentLevel+1, i == len(trace.Children)-1, prefix) // Recursively print children
	}
}

func formatInput(input string) string {
	if len(input) <= 9 {
		// If the input is 9 characters or fewer, just return it as is.
		return input
	}

	return input[:8] + "..." + input[len(input)-2:]
}

func hexToEther(hexValueStr string) string {
	// Check if hexValueStr is empty or invalid
	if hexValueStr == "" || (len(hexValueStr) > 2 && hexValueStr[:2] != "0x") {
		return "0 ETH"
	}

	// Remove leading "0x" if present
	if len(hexValueStr) > 2 {
		hexValueStr = hexValueStr[2:]
	}

	// Use big.Int for accuracy throughout
	hexValueBig, ok := new(big.Int).SetString(hexValueStr, 16)
	if !ok {
		log.Crit("hexToEther - failed to format tx hex Ether value -", "value", hexValueStr)
	}

	// Check if the value is zero before performing division
	if hexValueBig.Cmp(big.NewInt(0)) == 0 {
		return "0 ETH"
	}

	// Use big.Rat for precise division, then convert to string with formatting
	Ether := new(big.Rat).SetInt(big.NewInt(1e18))
	valueInEther := new(big.Rat).Quo(new(big.Rat).SetInt(hexValueBig), Ether)
	return valueInEther.FloatString(9) + " ETH"
}
