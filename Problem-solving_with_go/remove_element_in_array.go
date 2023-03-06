package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"
)

// TransactionData contains the properties of a transaction
type TransactionData struct {
	Sender   string
	Receiver string
	Amount   int64
	Time     int64
}

// Block contains the properties of a block in the blockchain
type Block struct {
	Timestamp     int64
	Transaction   TransactionData
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock creates a new block with the given transaction and previous block hash
func NewBlock(transaction TransactionData, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transaction, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates a new genesis block with a default transaction
func NewGenesisBlock() *Block {
	return NewBlock(TransactionData{"", "", 0, time.Now().Unix()}, []byte{})
}

// Serialize serializes the block into a byte slice
func (b *Block) Serialize() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

// HashTransaction hashes the transaction data and returns the resulting hash
func (t *TransactionData) HashTransaction() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(t)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())
	return hash[:]
}

// BlockChain is a chain of blocks
type BlockChain struct {
	Blocks []*Block
}

// AddBlock adds a new block to the blockchain
func (bc *BlockChain) AddBlock(transaction TransactionData) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transaction, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockChain creates a new blockchain with a genesis block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// ShowBlock shows the properties of a block
func ShowBlock(block *Block) {
	fmt.Printf("Timestamp: %d\n", block.Timestamp)
	fmt.Printf("Transaction: %s sent %d to %s\n", block.Transaction.Sender, block.Transaction.Amount, block.Transaction.Receiver)
	fmt.Printf("Prev. Block Hash: %x\n", block.PrevBlockHash)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Nonce: %d\n", block.Nonce)
}

// ShowBlockchain shows the properties of the entire blockchain
func ShowBlockchain(bc *BlockChain) {
	for _, block := range bc.Blocks {
		fmt.Println("---Block---")
		ShowBlock(block)
	}
}

// ProofOfWork represents the proof-of-work algorithm
type ProofOfWork struct {
	Block  *Block
	Target []byte
}

// NewProofOfWork creates a new proof-of-work object for the given block
func NewProofOfWork(b *Block) *ProofOfWork {
	target := make([]byte, 32)
	for i := 0; i < 32; i++ {
		target[i] = 0xff
	}
	return &ProofOfWork{b, target}
}

// Run performs the proof-of-work algorithm
func (pow *ProofOfWork) Run() (int, []byte) {
	var nonce int
	var hash [32]byte
	block := pow.Block

	for {
		data := append(block.Serialize(), intToHex(nonce)...)
		hash = sha256.Sum256(data)

		if bytes.Compare(hash[:], pow.Target) == -1 {
			break
		}
		nonce++
	}

	return nonce, hash[:]
}

func intToHex(num int) []byte {
	return []byte(strconv.FormatInt(int64(num), 16))
}

func main() {
	bc := NewBlockChain()

	bc.AddBlock(TransactionData{"fahim", "razzak", 100, time.Now().Unix()})
	bc.AddBlock(TransactionData{"razu", "manik", 200, time.Now().Unix()})

	fmt.Println("---Blockchain---")
	ShowBlockchain(bc)
}
