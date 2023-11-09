package crypto

import (
	"github.com/strecthr/testify"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	prvKey := GeneratePrivateKey()

	testify.assert.Equal(t, len(prvKey.Bytes()), prvKeyLen)
}
