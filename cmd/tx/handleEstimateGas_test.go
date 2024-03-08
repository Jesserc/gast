package transaction

import (
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func TestTryEstimateGasGas_Integration(t *testing.T) {
	testCases := []struct {
		name          string
		rpcURL        string // Test input: RPC URL
		from          string // Other test inputs could be added here
		to            string
		data          string
		wei           uint64
		expectedError string // Expected error substring
		wantGas       uint64 // Expected gas estimate (0 for failure cases)
	}{
		{
			name:          "Whitespace in RPC URL",
			rpcURL:        " https://rpc.mevblocker.io ",
			expectedError: "failed to dial RPC client",
		},
		{
			name:          "Malformed RPC URL",
			rpcURL:        "https://rpc.mevblocker.i",
			expectedError: "no such host",
		},
		{
			name:    "Successful Estimation",
			rpcURL:  "https://rpc.mevblocker.io",
			from:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			to:      "0xbe0eb53f46cd790cd13851d5eff43d12404d33e8",
			data:    "Hello Ethereum!",
			wei:     params.GWei * 20,
			wantGas: 21000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gas, err := TryEstimateGas(tc.rpcURL, tc.from, tc.to, tc.data, tc.wei)

			if tc.expectedError != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
				require.Greater(t, gas, tc.wantGas)
			}
		})
	}
}
