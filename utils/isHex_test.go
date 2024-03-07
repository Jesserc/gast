package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsHexWithOrWithout0xPrefix(t *testing.T) {
	isHex := IsHexWithOrWithout0xPrefix("Gast")
	require.False(t, isHex)

	isHex = IsHexWithOrWithout0xPrefix("0xc0ffeebabe")
	require.True(t, isHex)

	isHex = IsHexWithOrWithout0xPrefix("ffff")
	require.True(t, isHex)
}
