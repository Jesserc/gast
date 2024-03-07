package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEthConversion(t *testing.T) {
	v, err := EthConversion(10e18, "eth")
	require.Nil(t, err, "error should be nil")
	require.Equal(t, "10", v, "10e18 Wei should equal 10 ETH")

	v, err = EthConversion(10e18, "gwei")
	require.Nil(t, err, "error should be nil")
	require.Equal(t, "10000000000", v, "10e18 Wei should equal 10000000000 Gwei")

	v, err = EthConversion(10e18, "wei")
	require.Nil(t, err, "error should be nil")
	require.Equal(t, "10000000000000000000", v, "10e18 Wei should equal 10000000000000000000 Wei")

	v, err = EthConversion(10e18, "invalid")
	require.NotNil(t, err, "error should not be nil")
	require.ErrorContains(t, err, "denomination not supported")
}
