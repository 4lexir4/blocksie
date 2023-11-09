package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	prvKey := GeneratePrivateKey()
	println(prvKey)

	assert.Equal(t, len(prvKey.Bytes()), prvKeyLen)
}
