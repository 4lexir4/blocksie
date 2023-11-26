package types

import (
	"testing"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/proto"
	"github.com/4lexir4/blocksie/util"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRootHash(t *testing.T) {
	block := util.RandomBlock()
	tx := &proto.Transaction{
		Version: 1,
	}
	block.Transactions = append(block.Transactions, tx)
	assert.True(t, VerifyRootHash(block))
	assert.Equal(t, 32, len(block.Header.RoorHahs))
}

func TestSignVerifyBlock(t *testing.T) {
	var (
		block  = util.RandomBlock()
		prvKey = crypto.GeneratePrivateKey()
		pubKey = prvKey.Public()
	)

	sig := SignBlock(prvKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))

	assert.Equal(t, block.PublicKey, pubKey.Bytes())
	assert.Equal(t, block.Signature, sig.Bytes())

	assert.True(t, VerifyBlock(block))

	invalidPrvKey := crypto.GeneratePrivateKey()
	block.PublicKey = invalidPrvKey.Public().Bytes()

	assert.False(t, VerifyBlock(block))

}

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}
