package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block ...
type Block struct {
	Hash         string
	PrevBlockHas string
	Data         string
}

func (b *Block) hasBlock() {
	hash := sha256.Sum256([]byte(b.PrevBlockHas + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// Blockchain ...
type Blockchain struct {
	Blocks []*Block
}

// NewBlock ...
func NewBlock(data, prevBlockHash string) *Block {
	block := &Block{
		Data:         data,
		PrevBlockHas: prevBlockHash,
	}
	block.hasBlock()

	return block
}

// AddBlock ...
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, block)

	return block
}

// NewBlockChain ...
func NewBlockChain() *Blockchain {
	blocks := []*Block{NewBlock("Genesis Block", "")}

	return &Blockchain{Blocks: blocks}

}
