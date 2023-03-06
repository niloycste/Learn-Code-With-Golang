package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain.
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  string
	Hash      string
}

// Blockchain represents the full blockchain.
type Blockchain struct {
	chain []*Block
}

// NewBlock creates a new block and returns it.
func NewBlock(data string, prevHash string) *Block {
	block := &Block{
		Data:      data,
		Timestamp: GetTimestamp(),
		PrevHash:  prevHash,
		Hash:      CalculateHash(data, prevHash),
	}

	return block
}

// AddBlock adds a block to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.chain[len(bc.chain)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.chain = append(bc.chain, newBlock)
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", "")
	blockchain := &Blockchain{
		chain: []*Block{genesisBlock},
	}

	return blockchain
}

// CalculateHash calculates the hash of a block.
func CalculateHash(data string, prevHash string) string {
	hashData := data + prevHash
	hash := sha256.Sum256([]byte(hashData))
	return hex.EncodeToString(hash[:])
}

// GetTimestamp returns the current timestamp.
func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")

	for _, block := range bc.chain {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}
