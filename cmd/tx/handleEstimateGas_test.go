package transaction

import (
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func TestEstGas_Integration_FailureDueToWhitespaceInRPCURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	gas, err := TryEstimateGas(
		"https://rpc.mevblocker.io ", // Intentional whitespace should cause an error
		"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		"0xbe0eb53f46cd790cd13851d5eff43d12404d33e8",
		"Hello Ethereum!",
		params.GWei*20, // 20 gwei
	)

	require.NotNil(t, err)
	require.Zero(t, gas)
	require.ErrorContains(t, err, "failed to dial RPC client")
}

func TestEstGas_Integration_FailureWithMalformedRPCURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	gas, err := TryEstimateGas(
		"https://rpc.mevblocker.i", // Omit `o` at the end, this should also cause an error
		"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		"0xbe0eb53f46cd790cd13851d5eff43d12404d33e8",
		"Hello Ethereum!",
		params.GWei*20, // 20 gwei
	)

	require.NotNil(t, err)
	require.Zero(t, gas)
	require.ErrorContains(t, err, "no such host")
}

func TestEstGas_Integration_SuccessfulEstimation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	gas, err := TryEstimateGas(
		"https://rpc.mevblocker.io",
		"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		"0xbe0eb53f46cd790cd13851d5eff43d12404d33e8",
		"Hello Ethereum!",
		params.GWei*20, // 20 gwei
	)

	require.Nil(t, err)
	require.Greater(t, gas, uint64(21000)) // should be greater than 21000 gas because of the data field
}
