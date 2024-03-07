package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNonce_Integration(t *testing.T) {
	testCases := []struct {
		name          string
		address       string
		rpcURL        string
		expectedError string // Expected error substring

	}{
		{
			name:          "Whitespace in RPC URL",
			address:       "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			rpcURL:        " https://sepolia.drpc.org ",
			expectedError: "failed to dial RPC client",
		},
		{
			name:          "Malformed RPC URL",
			address:       "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			rpcURL:        "https://sepolia.drpc.or",
			expectedError: "no such host",
		},
		{
			name:    "Successful retrieval",
			address: "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			rpcURL:  "https://sepolia.drpc.org",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nonce, nextNonce, err := GetNonce(tc.address, tc.rpcURL)
			if tc.expectedError != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expectedError)
				require.Zero(t, nonce)
				require.Zero(t, nextNonce)
			} else {
				require.NoError(t, err)
				require.Greater(t, nextNonce, nonce)
			}
		})
	}
}
