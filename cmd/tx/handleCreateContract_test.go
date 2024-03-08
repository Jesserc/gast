package transaction

import (
	"testing"

	"github.com/Jesserc/gast/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateContract_Integration(t *testing.T) {
	testCases := []struct {
		name       string
		rpcURL     string
		dir        string
		privateKey string
		gasLimit   uint64
		wei        uint64
		wantError  bool
		errorMsg   string
	}{
		{
			name:       "RPC URL Whitespace Failure",
			rpcURL:     " https://sepolia.drpc.org ",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   21000,
			wantError:  true,
			errorMsg:   "failed to dial RPC client",
		},
		{
			name:       "Bad RPC URL Failure",
			rpcURL:     "https://sepolia.drpc.or",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   21000,
			wantError:  true,
			errorMsg:   "no such host",
		},
		{
			name:       "Invalid Private Key Failure",
			rpcURL:     "https://sepolia.drpc.org",
			privateKey: "123",
			gasLimit:   21000,
			wantError:  true,
			errorMsg:   "failed to decode private key: hex string of odd length",
		},
		{
			name:       "Private Key Size Exceeds Limit Failure",
			rpcURL:     "https://sepolia.drpc.org",
			privateKey: "ffffff2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   21000,
			wantError:  true,
			errorMsg:   "failed to convert private key to ECDSA",
		},
		{
			name:       "Insufficient Gas Failure",
			rpcURL:     "https://sepolia.drpc.org",
			dir:        "../contracts/CurrentYear.sol",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   0,
			wantError:  true,
			errorMsg:   "failed to send transaction: INTERNAL_ERROR: IntrinsicGas",
		},
		// This last test case makes an RPC call and has a wait time of 30-35 seconds
		{
			name:       "Successful Contract Deployment",
			rpcURL:     "https://sepolia.drpc.org",
			dir:        "../contracts/CurrentYear.sol",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   200000,
			wantError:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, _ := utils.CompileSol(tc.dir)

			txReceipt, err := CreateContract(tc.rpcURL, code, tc.privateKey, tc.gasLimit, tc.wei)

			if tc.wantError {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errorMsg)
				require.Nil(t, txReceipt)
			} else {
				require.NoError(t, err)
				require.NotNil(t, *txReceipt)
			}

		})
	}
}
