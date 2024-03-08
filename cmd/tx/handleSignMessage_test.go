package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignETHMessage(t *testing.T) {
	testCases := []struct {
		name        string
		message     string
		privKey     string
		expectedMsg string
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "Successful message signing",
			message:     "Jesserc",
			privKey:     "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			expectedMsg: "{\n \t\"address\": \"0x571B102323C3b8B8Afb30619Ac1d36d85359fb84\",\n \t\"msg\": \"Jesserc\",\n \t\"sig\": \"0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b\",\n \t\"version\": \"2\"\n }",
			wantErr:     false,
		},
		{
			name:    "Failure due to invalid private key",
			message: "Jesserc",
			privKey: "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662gggg",
			wantErr: true,
			errMsg:  "failed to convert private key to ECDSA: invalid hex character 'g' in private key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			signedETHMessage, err := SignETHMessage(tc.message, tc.privKey)

			if tc.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errMsg)
				require.Equal(t, tc.expectedMsg, "")
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedMsg, signedETHMessage)
			}
		})
	}
}
