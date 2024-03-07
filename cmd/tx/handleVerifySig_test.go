package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifySig(t *testing.T) {
	// Test verifying a signature with correct parameters - expects verification to succeed
	signed, err := VerifySig(
		"0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b",
		"0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
		"Jesserc",
	)
	require.NoError(t, err, "all params are correct, there shouldn't be any error")
	require.True(t, signed)

	// Test verifying a signature with slightly altered message - expects verification to fail
	signed, err = VerifySig(
		"0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b",
		"0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
		"Jessercc", // extra `c`
	)
	require.NoError(t, err, "all params are correct, there shouldn't be any error")
	require.False(t, signed)

	// Test with invalid signature (extra characters) - expects an error due to signature decoding failure
	signed, err = VerifySig(
		"0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541bgggg", // extra four `g`
		"0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
		"Jesserc",
	)
	require.Error(t, err, "some params are not correct, there should be an error")
	require.ErrorContains(t, err, "failed to decode signature to bytes")
	require.False(t, signed)

	// Test with invalid signature (extra zeros) - expects an error due to public key recovery failure
	signed, err = VerifySig(
		"0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b0000", // extra four `0`
		"0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
		"Jesserc",
	)
	require.Error(t, err, "some params are not correct, there should be an error")
	require.ErrorContains(t, err, "failed to recover public key bytes")
	require.False(t, signed)
}
