package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTransaction_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	testCases := []struct {
		name          string
		rpcURL        string
		address       string
		data          string
		privateKey    string
		gasLimit      uint64
		wei           uint64
		wantError     bool
		errorMsg      string
		wantTxReceipt bool
	}{
		{
			name:          "RPC URL Whitespace Failure",
			rpcURL:        " https://sepolia.drpc.org ",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey:    "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      21000,
			wantError:     true,
			errorMsg:      "failed to dial RPC client",
			wantTxReceipt: false,
		},
		{
			name:          "Bad RPC URL Failure",
			rpcURL:        "https://sepolia.drpc.or",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey:    "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      21000,
			wantError:     true,
			errorMsg:      "no such host",
			wantTxReceipt: false,
		},
		{
			name:          "Invalid Private Key Failure",
			rpcURL:        "https://sepolia.drpc.org",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey:    "123",
			gasLimit:      21000,
			wantError:     true,
			errorMsg:      "failed to decode private key: hex string of odd length",
			wantTxReceipt: false,
		},
		{
			name:          "Private Key Size Exceeds Limit Failure",
			rpcURL:        "https://sepolia.drpc.org",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey:    "ffffff2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      21000,
			wantError:     true,
			errorMsg:      "failed to convert private key to ECDSA",
			wantTxReceipt: false,
		},
		{
			name:          "Insufficient Gas Failure",
			rpcURL:        "https://sepolia.drpc.org",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:          "0xffffffffffff",
			privateKey:    "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      0,
			wantError:     true,
			errorMsg:      "failed to send transaction: INTERNAL_ERROR: IntrinsicGas",
			wantTxReceipt: false,
		},
		{
			name:          "Success With Hex Data",
			rpcURL:        "https://sepolia.drpc.org",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:          "0xffff",
			privateKey:    "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      25000,
			wei:           0.5e9,
			wantError:     false,
			wantTxReceipt: true,
		},
		{
			name:          "Success Hex Data No Prefix",
			rpcURL:        "https://sepolia.drpc.org",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:          "ffff",
			privateKey:    "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      25000,
			wei:           0.5e9,
			wantError:     false,
			wantTxReceipt: true,
		},
		{
			name:          "Success With String Data",
			rpcURL:        "https://sepolia.drpc.org",
			address:       "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:          "Gast",
			privateKey:    "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:      25000,
			wei:           0.5e9,
			wantError:     false,
			wantTxReceipt: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			txReceiptUrl, err := SendTransaction(
				tc.rpcURL,
				tc.address,
				"",
				tc.privateKey,
				tc.gasLimit,
				tc.wei,
			)

			if tc.wantError {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errorMsg)
				require.Equal(t, "", txReceiptUrl)
			} else {
				require.NoError(t, err)
				if tc.wantTxReceipt {
					require.Contains(t, txReceiptUrl, "https://sepolia.etherscan.io")
				} else {
					require.Equal(t, "", txReceiptUrl)
				}
			}
		})
	}
}
