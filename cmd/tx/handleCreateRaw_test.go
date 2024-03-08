package transaction

import (
	"testing"

	"github.com/Jesserc/gast/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRawTransaction_Integration(t *testing.T) {
	testCases := []struct {
		name       string
		rpcURL     string
		to         string
		data       string
		privateKey string
		gasLimit   uint64
		wei        uint64
		wantError  bool
		errorMsg   string
	}{
		{
			name:      "RPC URL Whitespace Failure",
			rpcURL:    " https://sepolia.drpc.org ",
			to:        "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			wantError: true,
			errorMsg:  "failed to dial RPC client",
		},
		{
			name:      "Bad RPC URL Failure",
			rpcURL:    "https://sepolia.drpc.or",
			to:        "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			wantError: true,
			errorMsg:  "no such host",
		},
		{
			name:       "Invalid Private Key Failure",
			rpcURL:     "https://sepolia.drpc.org",
			to:         "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			privateKey: "123",
			wantError:  true,
			errorMsg:   "failed to decode private key: hex string of odd length",
		},
		{
			name:       "Success With Hex Data",
			rpcURL:     "https://sepolia.drpc.org",
			to:         "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:       "0xffff",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   25000,
			wei:        0.5e9,
			wantError:  false,
		},
		{
			name:       "Success Hex Data No Prefix",
			rpcURL:     "https://sepolia.drpc.org",
			to:         "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:       "ffff",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   25000,
			wei:        0.5e9,
			wantError:  false,
		},
		{
			name:       "Success With String Data",
			rpcURL:     "https://sepolia.drpc.org",
			to:         "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
			data:       "Gast",
			privateKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			gasLimit:   25000,
			wei:        0.5e9,
			wantError:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rawTransaction, err := CreateRawTransaction(tc.rpcURL, tc.to, tc.data, tc.privateKey, tc.gasLimit, tc.wei)

			if tc.wantError {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errorMsg)
				require.Empty(t, rawTransaction)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, rawTransaction)
				require.True(t, utils.IsHexWithOrWithout0xPrefix(rawTransaction), "raw transaction should be a hexadecimal")
			}
		})
	}
}
