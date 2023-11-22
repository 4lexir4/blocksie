package node

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/proto"
	"github.com/4lexir4/blocksie/types"
	"github.com/4lexir4/blocksie/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randomBlock(t *testing.T, chain *Chain) *proto.Block {
	prvKey := crypto.GeneratePrivateKey()
	b := util.RandomBlock()
	prvBlock, err := chain.GetBlockByHeight(chain.Height())
	require.Nil(t, err)
	b.Header.PrvHash = types.HashBlock(prvBlock)
	types.SignBlock(prvKey, b)
	return b
}

func TestNewChain(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore(), NewMemoryTXStore())
	assert.Equal(t, 0, chain.Height())

	_, err := chain.GetBlockByHeight(0)
	assert.Nil(t, err)
}

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore(), NewMemoryTXStore())
	for i := 0; i < 100; i++ {
		b := randomBlock(t, chain)
		require.Nil(t, chain.AddBlock(b))
		require.Equal(t, chain.Height(), i+1)
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore(), NewMemoryTXStore())

	for i := 0; i < 100; i++ {
		block := randomBlock(t, chain)
		blockHash := types.HashBlock(block)
		require.Nil(t, chain.AddBlock(block))
		fetchedBlock, err := chain.GetBlockByHash(blockHash)
		require.Nil(t, err)
		require.Equal(t, block, fetchedBlock)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i + 1)
		require.Nil(t, err)
		require.Equal(t, block, fetchedBlockByHeight)
	}
}

func TestAddBlockWithTx(t *testing.T) {
	var (
		chain     = NewChain(NewMemoryBlockStore(), NewMemoryTXStore())
		block     = randomBlock(t, chain)
		prvKey    = crypto.NewPrivateKeyFromSeedString(genesisSeed)
		recepient = crypto.GeneratePrivateKey().Public().Address().Bytes()
	)

	ftt, err := chain.txStore.Get("8a814ba5ec1811952953f24421ef1c216e3f990e88994cb581e2f4ffc9a9513e")
	assert.Nil(t, err)
	fmt.Println(ftt)

	inputs := []*proto.TxInput{
		{
			PrvHash:     types.HashTransaction(ftt),
			PrvOutIndex: 0,
			PublicKey:   prvKey.Public().Bytes(),
		},
	}
	outputs := []*proto.TxOutput{
		{
			Amount:  100,
			Address: recepient,
		},
		{
			Amount:  900,
			Address: prvKey.Public().Address().Bytes(),
		},
	}
	tx := &proto.Transaction{
		Version: 1,
		Inputs:  inputs,
		Outputs: outputs,
	}
	block.Transactions = append(block.Transactions, tx)
	require.Nil(t, chain.AddBlock(block))
	txHash := hex.EncodeToString(types.HashTransaction(tx))

	fetchedTx, err := chain.txStore.Get(txHash)
	assert.Nil(t, err)
	assert.Equal(t, tx, fetchedTx)
}
