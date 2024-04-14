package hex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsHexWithOrWithout0xPrefix(t *testing.T) {
	isHex := WithOrWithout0xPrefix("Gast")
	require.False(t, isHex)

	isHex = WithOrWithout0xPrefix("0xc0ffeebabe")
	require.True(t, isHex)

	isHex = WithOrWithout0xPrefix("ffff")
	require.True(t, isHex)
}
