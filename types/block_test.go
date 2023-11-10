package types

import (
	"testing"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/util"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	var (
		block  = util.RandomBlock()
		prvKey = crypto.GeneratePrivateKey()
		pubKey = prvKey.Public()
	)

	sig := SignBlock(prvKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))
}

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}
