package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompileSol(t *testing.T) {
	testCases := []struct {
		name             string
		dir              string
		expectedBytecode string
		wantError        bool
		errorMsg         string
		equalMsg         string
	}{
		{
			name:             "Failure Due To Empty Directory",
			dir:              "",
			expectedBytecode: "",
			wantError:        true,
			errorMsg:         "error should not be nil",
			equalMsg:         "compiled bytecode for empty dir should be empty",
		},
		{
			name:             "Successful Compilation",
			dir:              "../contracts/CurrentYear.sol",
			expectedBytecode: "608060405234801561000f575f80fd5b506107e85f81905550610143806100255f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063ef88a09214610038578063fd08921b14610054575b5f80fd5b610052600480360381019061004d91906100ba565b610072565b005b61005c61007b565b60405161006991906100f4565b60405180910390f35b805f8190555050565b5f8054905090565b5f80fd5b5f819050919050565b61009981610087565b81146100a3575f80fd5b50565b5f813590506100b481610090565b92915050565b5f602082840312156100cf576100ce610083565b5b5f6100dc848285016100a6565b91505092915050565b6100ee81610087565b82525050565b5f6020820190506101075f8301846100e5565b9291505056fea2646970667358221220676f9af6af517a2017a5135a8cf7b8a7cb55ac9fde9b3df5a229855fdfdc4c7b64736f6c63430008180033",
			wantError:        false,
			errorMsg:         "error should be nil",
			equalMsg:         "compiled bytecode for ../contracts/CurrentYear.sol should be equal to expected",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := CompileSol(tc.dir)
			if tc.wantError {
				require.Error(t, err, tc.errorMsg)
				require.Equal(t, tc.expectedBytecode, code, tc.equalMsg)
			} else {
				require.NoError(t, err, tc.errorMsg)
				require.Equal(t, tc.expectedBytecode, code, tc.equalMsg)
			}
		})
	}
}
