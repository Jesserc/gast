package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifySig(t *testing.T) {
	testCases := []struct {
		name    string
		sig     string
		address string
		message string
		wantErr bool
		errMsg  string
		signed  bool
	}{
		{
			name:    "Altered Message Verification Failure",
			sig:     "0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b",
			address: "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			message: "Jessercc",
			wantErr: false,
			signed:  false,
		},
		{
			name:    "Invalid Signature Extra Characters",
			sig:     "0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541bgggg",
			address: "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			message: "Jesserc",
			wantErr: true,
			errMsg:  "failed to decode signature to bytes",
			signed:  false,
		},
		{
			name:    "Invalid Signature Extra Zeros",
			sig:     "0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b0000",
			address: "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			message: "Jesserc",
			wantErr: true,
			errMsg:  "failed to recover public key bytes",
			signed:  false,
		},
		{
			name:    "Successful Verification",
			sig:     "0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b",
			address: "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			message: "Jesserc",
			wantErr: false,
			signed:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			signed, err := VerifySig(tc.sig, tc.address, tc.message)

			if tc.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errMsg)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.signed, signed) // ensure it's true for tests that expects true or false for those that expect false
		})
	}
}
