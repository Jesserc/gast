package gasprice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGasPrice_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	gasPrice, err := GetCurrentGasPrice(" https://rpc.mevblocker.io ") // Intentional whitespace
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to dial RPC client")
	require.Zero(t, gasPrice)

	gasPrice, err = GetCurrentGasPrice("https://rpc.mevblocker.io")
	require.NoError(t, err)
	require.NotZero(t, gasPrice)
}
