package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendRawTransaction_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	
	testCases := []struct {
		name      string
		rpcURL    string
		rawTx     string
		wantError bool
		errorMsg  string
	}{
		{
			name:      "RPC URL Whitespace Failure",
			rpcURL:    " https://sepolia.drpc.org ",
			wantError: true,
			errorMsg:  "failed to dial RPC client",
		},
		{
			name:      "Invalid Raw Transaction hex",
			rpcURL:    "https://sepolia.drpc.or",
			rawTx:     "gggg",
			wantError: true,
			errorMsg:  "failed to decode raw transaction to rlp decoded bytes",
		},
		{
			name:      "Bad RPC URL Failure",
			rawTx:     "b87502f87283aa36a7820156843536a9a184e99e2f66825208944924fb92285cb10bc440e6fb4a53c2b94f2930c58398968080c001a04d7e76a9b1f3dc45f0f57b93e39842f1ee8c03f007a10363b2d3ddc8084ca00ca078f9f2a1c311db686c727aef6f03670e7b8100caf2368f7e7bcd5ce69d65c7e8",
			rpcURL:    "https://sepolia.drpc.or",
			wantError: true,
			errorMsg:  "no such host",
		},
		{
			name:      "Invalid Raw TX length",
			rawTx:     "b87502f87283aa36a7820156843536a9a184e99e2f66825208944924fb92285cb10bc440e6fb4a53c2b94f2930c58398968080c001a04d7e76a9b1f3dc45f0f57b93e39842f1ee8c03f007a10363b2d3ddc8084ca00ca078f9f2a1c311db686c727aef6f03670e7b8100caf2368f7e7bcd5ce69d65c7e8ffff", // Invalid `ffff`
			rpcURL:    "https://sepolia.drpc.org",
			wantError: true,
			errorMsg:  "failed to decode transaction rlp bytes to types.Transaction",
		},
		// Commented out because the raw tx has been sent already. Replace this with a new one and run test
		// {
		// 	name:      "Successful Raw Transaction Propagation",
		// 	rawTx:     "b87502f87283aa36a7820156843536a9a184e99e2f66825208944924fb92285cb10bc440e6fb4a53c2b94f2930c58398968080c001a04d7e76a9b1f3dc45f0f57b93e39842f1ee8c03f007a10363b2d3ddc8084ca00ca078f9f2a1c311db686c727aef6f03670e7b8100caf2368f7e7bcd5ce69d65c7e8",
		// 	rpcURL:    "https://sepolia.drpc.org",
		// 	wantError: false,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			txReceiptUrl, txJson, err := SendRawTransaction(
				tc.rawTx,
				tc.rpcURL,
			)

			if tc.wantError {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errorMsg)
				require.Equal(t, "", txReceiptUrl)
				require.Empty(t, txJson)
			} else {
				require.NoError(t, err)
				require.Contains(t, txReceiptUrl, "https://sepolia.etherscan.io/tx/")
			}
		})
	}
}
