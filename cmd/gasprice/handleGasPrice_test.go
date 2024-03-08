package gasprice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGasPrice_IntegrationTDD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	
	testCases := []struct {
		name      string
		rpcURL    string
		wantError bool
		errorMsg  string
	}{
		{
			name:      "RPC URL Whitespace Failure",
			rpcURL:    " https://rpc.mevblocker.io ",
			wantError: true,
			errorMsg:  "failed to dial RPC client",
		},
		{
			name:   "Successful Gas Price Retrieval",
			rpcURL: "https://rpc.mevblocker.io",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gasPrice, err := GetCurrentGasPrice(tc.rpcURL)
			if tc.wantError {
				require.Error(t, err)
				require.Zero(t, gasPrice)
				require.ErrorContains(t, err, tc.errorMsg)
			} else {
				require.NoError(t, err)
				require.NotZero(t, gasPrice)
			}
		})
	}
}
