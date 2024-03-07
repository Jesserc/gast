package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTx_Integration_RPCURLWhitespaceFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		" https://sepolia.drpc.org ", // Intentional whitespace should cause an error
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"",
		"2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
		21000,
		0.5e9, // 0.5 gwei
	)

	require.Error(t, err)
	require.ErrorContains(t, err, "failed to dial RPC client")
	require.Equal(t, txReceiptUrl, "") // tx receipt url should be an empty string
}

func TestSendTx_Integration_BadRPCURLFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.or", // Omit `g` in RPC URL
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"",
		"2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
		21000,
		0.5e9, // 0.5 gwei
	)

	require.Error(t, err)
	require.ErrorContains(t, err, "failed to get chain ID")
	require.Equal(t, txReceiptUrl, "")
}

func TestSendTx_Integration_InvalidPrivateKeyFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.org",
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"",
		"123", // Invalid private
		21000,
		0.5e9, // 0.5 gwei
	)

	require.Error(t, err)
	require.ErrorContains(t, err, "failed to decode private key: hex string of odd length")
	require.Equal(t, txReceiptUrl, "")
}

func TestSendTx_Integration_PrivateKeySizeExceedsLimitFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.org",
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"",
		"ffffff2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662", // Invalid private key: exceeds 256-bit size limit
		21000,
		0.5e9, // 0.5 gwei
	)

	require.Error(t, err)
	require.ErrorContains(t, err, "failed to convert private key to ECDSA")
	require.Equal(t, txReceiptUrl, "")
}

func TestSendTx_Integration_InsufficientGasFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.org",
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"0xffff", // hex data
		"2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
		21000, // use less than required gas to cause an error
		0.5e9,
	)

	require.Error(t, err, "returned error should not be nil")
	require.ErrorContains(t, err, "failed to send transaction: INTERNAL_ERROR: IntrinsicGas")
	require.Equal(t, txReceiptUrl, "", "transaction receipt url should be empty")
}

func TestSendTx_Integration_SuccessWithHexData(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.org",
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"0xffff", // hex data
		"2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
		25000,
		0.5e9,
	)

	require.NoError(t, err, "returned error should be nil")
	require.Contains(t, txReceiptUrl, "https://sepolia.etherscan.io")
}

func TestSendTx_Integration_SuccessHexDataNoPrefix(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.org",
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"ffff", // hex data without `0x` prefix
		"2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
		25000,
		0.5e9,
	)

	require.NoError(t, err, "returned error should be nil")
	require.Contains(t, txReceiptUrl, "https://sepolia.etherscan.io")
}

func TestSendTx_Integration_SuccessWithStringData(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	txReceiptUrl, err := SendTransaction(
		"https://sepolia.drpc.org",
		"0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5",
		"Gast", // with normal string data
		"2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
		25000,
		0.5e9,
	)

	require.NoError(t, err, "returned error should be nil")
	require.Contains(t, txReceiptUrl, "https://sepolia.etherscan.io")
}
