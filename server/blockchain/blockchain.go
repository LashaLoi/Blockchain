package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block is item of blockchain
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

func (b *Block) hasBlock() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// Blockchain is a list of blocks
type Blockchain struct {
	Blocks []*Block
}

// NewBlock allow you to create a new block
func NewBlock(data, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.hasBlock()

	return block
}

// AddBlock allow you tp add block in blockchain
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, block)

	return block
}

// NewBlockChain allow you to create a new blockchain
func NewBlockChain() *Blockchain {
	blocks := []*Block{NewBlock("Genesis Block", "")}

	return &Blockchain{Blocks: blocks}

}
