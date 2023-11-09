package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	prvKey := GeneratePrivateKey()
	assert.Equal(t, len(prvKey.Bytes()), prvKeyLen)

	pubKey := prvKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TerstPrivateKeySign(t *testing.T) {
	prvKey := GeneratePrivateKey()
	pubKey := prvKey.Public()
	msg := []byte("nana lala dada")

	sig := prvKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))

	// Test with invalid message
	assert.False(t, sig.Verify(pubKey, []byte("haha")))

	// Test with not matching publick key
	nonMatchingPrvKey := GeneratePrivateKey()
	nonMatchingPubKey := nonMatchingPrvKey.Public()
	assert.False(t, sig.Verify(nonMatchingPubKey, msg))

}

func TestPublicKeyToAddress(t *testing.T) {
	prvKey := GeneratePrivateKey()
	pubKey := prvKey.Public()
	address := pubKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
}
