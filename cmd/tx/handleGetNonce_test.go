package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNonceIntegration_FailureDueToWhitespaceInRPCURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	address1 := "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84"
	rpcUrl := " https://sepolia.drpc.org " // Intentional whitespace should cause an error

	cNonce, nNonce, err := GetNonce(address1, rpcUrl)

	require.NotNil(t, err)
	require.ErrorContains(t, err, "failed to dial RPC client")

	require.Zero(t, cNonce, "current nonce should be zero for invalid rpc url")
	require.Zero(t, nNonce, "next nonce should be zero for invalid rpc url")
}

func TestGetNonceIntegration_FailureWithMalformedRPCURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	address1 := "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84"
	rpcUrl := "https://sepolia.drpc.or" // Omit `o` at the end, this should also cause an error

	cNonce, nNonce, err := GetNonce(address1, rpcUrl)

	require.NotNil(t, err)
	require.ErrorContains(t, err, "no such host")

	require.Zero(t, cNonce, "current nonce should be zero for invalid rpc url")
	require.Zero(t, nNonce, "next nonce should be zero for invalid rpc url")
}

func TestGetNonceIntegration_SuccessfulRetrieval(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	address1 := "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84"
	rpcUrl := "https://sepolia.drpc.org"

	cNonce, nNonce, err := GetNonce(address1, rpcUrl)

	require.Nil(t, err)
	require.Greater(t, nNonce, cNonce)
}
