package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendBlobTX_Integration(t *testing.T) {
	testCases := []struct {
		name        string
		rpcURL      string
		address     string
		data        string
		privateKey  string
		saveBlobDir string
		wantError   bool
		errorMsg    string
		wantTxHash  bool
	}{
		{
			name:       "RPC URL Whitespace Failure",
			rpcURL:     " https://sepolia.drpc.org ",
			address:    "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			wantError:  true,
			errorMsg:   "failed to dial RPC client",
		},
		{
			name:       "Bad RPC URL Failure",
			rpcURL:     "https://sepolia.drpc.or",
			address:    "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			wantError:  true,
			errorMsg:   "no such host",
		},
		{
			name:       "Invalid Private Key Failure",
			rpcURL:     "https://sepolia.drpc.org",
			address:    "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey: "123",
			wantError:  true,
			errorMsg:   "failed to decode private key: hex string of odd length",
		},
		{
			name:       "Private Key Size Exceeds Limit Failure",
			rpcURL:     "https://sepolia.drpc.org",
			address:    "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey: "ffffff2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			wantError:  true,
			errorMsg:   "failed to convert private key to ECDSA",
		},
		{
			name:        "Compute Blob Commitment Failure",
			rpcURL:      "https://sepolia.drpc.org",
			address:     "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:        "ffff",
			privateKey:  "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			saveBlobDir: "gast/blob-tx",
			wantError:   true,
			errorMsg:    "failed to compute blob commitment: scalar is not canonical when interpreted as a big integer in big-endian",
		},
		{
			name:        "Success With String Data",
			rpcURL:      "https://sepolia.drpc.org",
			address:     "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:        "Hello Blobs!",
			privateKey:  "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			saveBlobDir: "gast/blob-tx",
			wantError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			txHash, err := SendBlobTX(
				tc.rpcURL,
				tc.address,
				tc.data,
				tc.privateKey,
				tc.saveBlobDir,
			)

			if tc.wantError {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errorMsg)
				require.Equal(t, "", txHash)
			} else {
				require.NoError(t, err)
				require.Equal(t, len(txHash), 66)
				require.NotEmpty(t, txHash)
			}
		})
	}
}
