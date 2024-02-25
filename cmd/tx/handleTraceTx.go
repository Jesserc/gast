package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Jesserc/gast/cmd/tx/params"
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
	var (
		client *rpc.Client
		err    error
	)

	if rpcUrl == "" {
		client, err = rpc.Dial("https://rpc.builder0x69.io/")
		if err != nil {
			return "", err
		}
	} else {
		client, err = rpc.Dial(rpcUrl)
		if err != nil {
			return "", err
		}
	}

	var result json.RawMessage
	err = client.CallContext(context.Background(), &result, "ots_traceTransaction", hash)
	if err != nil {
		return "", err
	}

	var traces []Trace
	if err := json.Unmarshal(result, &traces); err != nil {
		return "", err
	}

	rootTrace := buildTraceHierarchy(traces)

	printTrace(rootTrace, 0, false, "")
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

func printTrace(trace *Trace, indentLevel int, isLastChild bool, prefix string) {
	var indent, currentPrefix string
	if indentLevel > 0 {
		indent = strings.Repeat("    ", indentLevel-1) // Basic indentation for hierarchy level
		if isLastChild {
			currentPrefix = prefix + "└── " + params.ColorGreen + "← " + params.ColorReset
			prefix += "    " // Extend the prefix for child traces without a connecting line
		} else {
			currentPrefix = prefix + "├── " + params.ColorGreen + "← " + params.ColorReset
			prefix += "│   " // Add a vertical line for child traces
		}
	}
	formattedInput := formatInput(trace.Input) // Format the input field
	fmt.Printf("%s%s%sType:%s %s, %sFrom:%s %s, %sTo:%s %s, %sDepth:%s %d, %sInput:%s [%s]\n",
		indent, currentPrefix,
		params.ColorGreen, params.ColorReset,
		trace.Type, params.ColorGreen,
		params.ColorReset, trace.From,
		params.ColorGreen, params.ColorReset,
		trace.To, params.ColorGreen,
		params.ColorReset, trace.Depth,
		params.ColorGreen, params.ColorReset,
		formattedInput,
	)

	for i, child := range trace.Children {
		printTrace(child, indentLevel+1, i == len(trace.Children)-1, prefix) // Recursively print children
	}
}
func formatInput(input string) string {
	if len(input) <= 4 {
		// If the input is 4 characters or fewer, just return it as is.
		return input
	}

	return input[:8] + "..." + input[len(input)-4:]
}
