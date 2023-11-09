package crypto

import (
	//  Needed to generate the seed below
	//"crypto/rand"
	//"encoding/hex"
	//"fmt"
	//"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	prvKey := GeneratePrivateKey()
	assert.Equal(t, len(prvKey.Bytes()), prvKeyLen)

	pubKey := prvKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "1053166da15d6da372a520d99a631e7bb4dfcb1d9089b458fd25ce93387dd4bb"
		prvKey     = NewPrivateKeyFromString(seed)
		addressStr = "e911ea8baa03eca3d4009ee7ec334fa1a42a31ac"
	)

	// How the seed above was generated:
	//seed := make([]byte, 32)
	//io.ReadFull(rand.Reader, seed)
	//fmt.Println(hex.EncodeToString(seed))

	assert.Equal(t, prvKeyLen, len(prvKey.Bytes()))

	// How to get address string:
	//address := prvKey.Public().Address()
	//fmt.Println(address)

	address := prvKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
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
