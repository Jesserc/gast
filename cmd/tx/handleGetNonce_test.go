package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleGetNonceFailureOne(t *testing.T) {
	address1 := "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84"
	rpcUrl := " `https://sepolia.drpc.org`" // Intentional whitespace should cause an error

	cNonce, nNonce, err := GetNonce(address1, rpcUrl)

	require.NotNil(t, err)
	require.Contains(t, err.Error(), "failed to dial RPC client")

	require.Zero(t, cNonce, "current nonce should be zero for invalid rpc url")
	require.Zero(t, nNonce, "next nonce should be zero for invalid rpc url")
}

func TestHandleGetNonceFailureTwo(t *testing.T) {
	address1 := "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84"
	rpcUrl := "https://sepolia.drpc.or" // Omit `g` at the end

	cNonce, nNonce, err := GetNonce(address1, rpcUrl)

	require.NotNil(t, err)
	require.Contains(t, err.Error(), "failed to get nonce")

	require.Zero(t, cNonce, "current nonce should be zero for invalid rpc url")
	require.Zero(t, nNonce, "next nonce should be zero for invalid rpc url")
}

func TestHandleGetNonce(t *testing.T) {
	address1 := "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84"
	rpcUrl := "https://sepolia.drpc.org"

	cNonce, nNonce, err := GetNonce(address1, rpcUrl)

	require.Nil(t, err)

	require.Greater(t, nNonce, cNonce)
}
