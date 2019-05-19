package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Blockchain is our digital spine
type Blockchain struct {
	blocks []*Block
}

// Block is each singular unit in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash calculates a hash value for a Block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	fmt.Println("info: ", info)
	hash := sha256.Sum256(info)
	fmt.Println("hash: ", hash)
	b.Hash = hash[:]
}

// CreateBlock returns a new block hashed from the latest block in the chain
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash}
	block.DeriveHash()
	return block
}

// AddBlock will add a new block to the tail of its chain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis creates our first block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockchain creates our blockchain with its new block
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	fmt.Println("A blockchain written in Go")
	chain := InitBlockchain()

	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	chain.AddBlock("Fourth block")

	fmt.Println("chain: ", chain)

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
