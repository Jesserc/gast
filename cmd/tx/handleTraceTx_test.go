package transaction

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTraceTxIntegration_FailureDueToWhitespaceInRPCURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	trace, err := TraceTx(
		"0xd12e31c3274ff32d5a73cc59e8deacbb0f7ac4c095385add3caa2c52d01164c1",
		"  https://rpc.builder0x69.io/ ", // Intentional whitespace to cause an error
	)

	require.Nil(t, trace)
	require.NotNil(t, err)

	require.ErrorContains(t, err, "failed to dial RPC client")
}

func TestTraceTxIntegration_FailureDueToInvalidTransactionHash(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	trace, err := TraceTx(
		"0xd12e31c3274ff32d5a73cc", // Invalid hash
		"https://rpc.builder0x69.io/",
	)

	require.Nil(t, trace)
	require.NotNil(t, err)

	require.ErrorContains(t, err, "invalid argument")
}

func TestTraceTxIntegration_SuccessWithDefaultRPCURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	trace, err := TraceTx(
		"0xd12e31c3274ff32d5a73cc59e8deacbb0f7ac4c095385add3caa2c52d01164c1", // invalid hash
		"", // Use default RPC URL by leaving this empty
	)

	require.NotNil(t, trace)
	require.Nil(t, err)

	require.Equal(t, reflect.TypeOf(trace), reflect.TypeOf(&Trace{}))
}

func TestTraceTxIntegration_SuccessWithSpecifiedRPCURL(t *testing.T) {
	trace, err := TraceTx(
		"0xd12e31c3274ff32d5a73cc59e8deacbb0f7ac4c095385add3caa2c52d01164c1",
		"https://rpc.builder0x69.io/",
	)

	require.NotNil(t, trace)
	require.Nil(t, err)

	require.Equal(t, reflect.TypeOf(trace), reflect.TypeOf(&Trace{}))

}

func TestTraceTxIntegration_PrintTrace(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	trace, err := TraceTx(
		"0xd12e31c3274ff32d5a73cc59e8deacbb0f7ac4c095385add3caa2c52d01164c1", // Example hash
		"", // Use default RPC URL by leaving this empty
	)

	require.NotNil(t, trace, "Trace should not be nil.")
	require.Nil(t, err, "Should not error while tracing transaction.")

	require.Equal(t, reflect.TypeOf(trace), reflect.TypeOf(&Trace{}), "Expected a *Trace type.")

	// Save the original os.Stdout to restore it later.
	originalStdout := os.Stdout

	// Create a pipe: 'r' for reading, 'w' for writing.
	// This allows capturing the output directed to os.Stdout.
	r, w, _ := os.Pipe()
	// Temporarily set os.Stdout to our pipe's writer.
	os.Stdout = w

	printTrace(trace, 0, false, "") // printing trace to stdout is now redirected to our pipe `w`
	w.Close()                       // Close the writer end of the pipe to signal we're done writing.

	// Read all output captured from the pipe.
	out, _ := io.ReadAll(r)
	os.Stdout = originalStdout // Restore the original stdout to avoid side effects on test outputs

	require.Contains(t, string(out), "\033[32mType:\033[0m CALL, \033[32mFrom:\033[0m 0x734bce0ca8f39c2f9768267390adf7df0d615db7, \033[32mTo:\033[0m 0x3fc91a3afd70395cd496c647d5a6cc9d4b2b7fad, \033[32mDepth:\033[0m 0, \033[32mValue:\033[0m 0 ETH, \033[32mInput:\033[0m [0x359356...00]\n", "Output should contain trace information.")
}

func TestTraceTx_FormatInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"0xc0ffebabe", "0xc0ffeb...be"},
		{"", ""},
		{"0xdeAdBeef", "0xdeAdBe...ef"},
		{"0xbeef", "0xbeef"},
	}

	for _, c := range cases {
		out := formatInput(c.input)
		require.Equal(t, c.expected, out)
	}
}

func TestTraceTx_HexToEtherConversion(t *testing.T) {
	cases := []struct {
		hexValueStr string
		expected    string
	}{
		{"0x0", "0 ETH"},
		{"0xde0b6b3a7640000", "1.000000000 ETH"}, // 10^18 in hex
		{"0xdead0b6b3a7640000", "256.727897641 ETH"},
		{"", "0 ETH"},
	}

	for _, c := range cases {
		eth := hexToEther(c.hexValueStr)
		require.Equal(t, c.expected, eth)
	}
}
