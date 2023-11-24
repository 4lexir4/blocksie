package node

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/proto"
	"github.com/4lexir4/blocksie/types"
)

const genesisSeed = "ec82fa001561cdf9ee76af621e5ef8f995dd1f077cd22fe56a929e06967aabf9"

type HeaderList struct {
	headers []*proto.Header
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		headers: []*proto.Header{},
	}
}

func (list *HeaderList) Add(h *proto.Header) {
	list.headers = append(list.headers, h)
}

func (list *HeaderList) Get(index int) *proto.Header {
	if index > list.Height() {
		panic("Index too hight.")
	}

	return list.headers[index]
}

func (list *HeaderList) Height() int {
	return list.Len() - 1
}

func (list *HeaderList) Len() int {
	return len(list.headers)
}

type UTXO struct {
	Hash     string
	OutIndex int
	Amount   int64
	Spent    bool
}

type Chain struct {
	txStore    TXStorer
	blockStore BlockStorer
	utxoStore  UTXOStorer
	headers    *HeaderList
}

func NewChain(bs BlockStorer, txStore TXStorer) *Chain {
	chain := &Chain{
		blockStore: bs,
		txStore:    txStore,
		utxoStore:  NewMemoryUTXOStore(),
		headers:    NewHeaderList(),
	}

	chain.addBlock(createGenesisBlock())
	return chain
}

func (c *Chain) Height() int {
	return c.headers.Height()
}

func (c *Chain) AddBlock(b *proto.Block) error {
	if err := c.ValidateBlock(b); err != nil {
		return err
	}
	return c.addBlock(b)
}

func (c *Chain) addBlock(b *proto.Block) error {
	// add the header to the list of headers
	c.headers.Add(b.Header)

	for _, tx := range b.Transactions {
		if err := c.txStore.Put(tx); err != nil {
			return err
		}

		hash := hex.EncodeToString(types.HashTransaction(tx))

		for it, output := range tx.Outputs {
			utxo := &UTXO{
				Hash:     hash,
				Amount:   output.Amount,
				OutIndex: it,
				Spent:    false,
			}
			if err := c.utxoStore.Put(utxo); err != nil {
				return err
			}
		}
	}
	// validation
	return c.blockStore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.blockStore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if c.Height() < height {
		return nil, fmt.Errorf("Given height [%d] too high - height is [%d].", height, c.Height())
	}

	header := c.headers.Get(height)
	hash := types.HashHeader(header)
	return c.GetBlockByHash(hash)

}

func (c *Chain) ValidateBlock(b *proto.Block) error {
	// Validate the signature of the block
	if !types.VerifyBlock(b) {
		return fmt.Errorf("Invalid block signature")
	}

	// Validate if the prvHash is the actual hash of the current clock
	currentBlock, err := c.GetBlockByHeight(c.Height())
	if err != nil {
		return err
	}

	hash := types.HashBlock(currentBlock)
	if !bytes.Equal(hash, b.Header.PrvHash) {
		return fmt.Errorf("Invalid previous block hash")
	}

	for _, tx := range b.Transactions {
		if err := c.ValidateTransaction(tx); err != nil {
			return err
		}

	}

	return nil
}

func (c *Chain) ValidateTransaction(tx *proto.Transaction) error {
	// verify the signature
	if !types.VerifyTransaction(tx) {
		return fmt.Errorf("Invalid tx signatrue")
	}
	// check if all the inputs are unspent
	nInputs := len(tx.Inputs)
	hash := hex.EncodeToString(types.HashTransaction(tx))
	sumInputs := 0
	for i := 0; i < nInputs; i++ {
		var (
			prvHash   = hex.EncodeToString(tx.Inputs[i].PrvHash)
			key       = fmt.Sprintf("%s_%d", prvHash, i)
			utxo, err = c.utxoStore.Get(key)
		)
		sumInputs += int(utxo.Amount)
		if err != nil {
			return err
		}
		if utxo.Spent {
			return fmt.Errorf("Input %d of tx %s is already spent", i, hash)
		}
	}
	sumOutputs := 0
	for _, output := range tx.Outputs {
		sumOutputs += int(output.Amount)
	}

	if sumInputs < sumOutputs {
		return fmt.Errorf("Insufficient balance, got [%d], need [%d]", sumInputs, sumOutputs)
	}

	return nil
}

func createGenesisBlock() *proto.Block {
	prvKey := crypto.NewPrivateKeyFromSeedString(genesisSeed)
	block := &proto.Block{
		Header: &proto.Header{
			Version: 1,
		},
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{},
		Outputs: []*proto.TxOutput{
			{
				Amount:  1000,
				Address: prvKey.Public().Address().Bytes(),
			},
		},
	}

	block.Transactions = append(block.Transactions, tx)
	types.SignBlock(prvKey, block)

	return block
}
