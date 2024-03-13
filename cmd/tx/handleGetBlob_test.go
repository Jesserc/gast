package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBlob(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	testCases := []struct {
		name                  string
		blockRootOrSlotNumber string
		kzgCommitment         string
		expectedError         string
		isNull                bool
	}{
		{
			name:                  "Invalid slot number",
			blockRootOrSlotNumber: "###",
			kzgCommitment:         "0xb28e4d255047f6e50b3d7548d37155b6e2289e82520aa6248d9fbe50e73b81d9f705cb3f2192d55caf54e26fb29c419a",
			expectedError:         "GET failed with status 404",
		},
		{
			name:                  "Invalid block root",
			blockRootOrSlotNumber: "0x6895707e38c30605b4d76cda082fd0173a9ef99686e747fc673e0923acf0acbd",
			kzgCommitment:         "0xb28e4d255047f6e50b3d7548d37155b6e2289e82520aa6248d9fbe50e73b81d9f705cb3f2192d55caf54e26fb29c419a",
			expectedError:         "NOT_FOUND: beacon block with root 0x6895…acbd",
		},
		{
			name:                  "Invalid kzg commitment",
			blockRootOrSlotNumber: "8626178",
			kzgCommitment:         "$$$",
			expectedError:         "failed to decode hex kzg commitment: hex string without 0x prefix",
		},
		{
			name:                  "Success with valid block root",
			blockRootOrSlotNumber: "0x2eece2eff327e0f611672169ebffb9c1cf3085433f98ca445ba1011507e90d69",
			kzgCommitment:         "0xb28e4d255047f6e50b3d7548d37155b6e2289e82520aa6248d9fbe50e73b81d9f705cb3f2192d55caf54e26fb29c419a",
		},
		{
			name:                  "Success with valid slot number",
			blockRootOrSlotNumber: "8626178",
			kzgCommitment:         "0xb28e4d255047f6e50b3d7548d37155b6e2289e82520aa6248d9fbe50e73b81d9f705cb3f2192d55caf54e26fb29c419a",
		},
		{
			name:                  "Empty result",
			blockRootOrSlotNumber: "8626178",
			kzgCommitment:         "0xb28e4d255047f6e50b3d7548d37155b6e2289e82520aa6248d9fbe50e73b81d9f705cb3f2192d55caf54e26fb29c419a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			blob, err := GetBlob(tc.blockRootOrSlotNumber, tc.kzgCommitment)
			if tc.expectedError != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expectedError)
				require.Empty(t, blob)
			} else {
				if tc.isNull {
					require.NoError(t, err)
					require.Equal(t, "null", blob)
				}
				require.NoError(t, err)
				require.NotEmpty(t, blob)
			}
		})
	}
}
