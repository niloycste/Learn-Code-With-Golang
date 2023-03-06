package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp  int64
	Data       []byte
	Hash       []byte
	PrevHash   []byte
	Nonce      int
	Difficulty int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	nonce := []byte(strconv.Itoa(b.Nonce))
	difficulty := []byte(strconv.Itoa(b.Difficulty))
	blockData := append(b.Data, b.PrevHash...)
	blockData = append(blockData, timestamp...)
	blockData = append(blockData, nonce...)
	blockData = append(blockData, difficulty...)
	hash := sha256.Sum256(blockData)
	b.Hash = hash[:]
}

func (b *Block) GetData() []byte {
	return b.Data
}

func NewBlock(data []byte, prevHash []byte, nonce int, difficulty int, timestamp int64) *Block {
	block := &Block{timestamp, data, []byte{}, prevHash, nonce, difficulty}
	block.SetHash()
	return block
}

func FetchDataFromBlocks(block1, block2 *Block) ([]byte, int64, int, int) {
	data := append(block1.Data, block2.Data...)
	timestamp := time.Now().Unix()
	nonce := block1.Nonce + block2.Nonce
	difficulty := block1.Difficulty + block2.Difficulty
	return data, timestamp, nonce, difficulty
}

func main() {
	block1 := NewBlock([]byte("Hemel Hasan Transafer 100$ to Niloy's Wallet and "), []byte{}, 12, 3, time.Now().Unix())
	block2 := NewBlock([]byte("Niloy Islam Has Received 100$ from Hemel Hasan"), block1.Hash, 15, 4, time.Now().Unix())
	fmt.Println("Genesis Block  Timestamp: ", block1.Timestamp)
	fmt.Println("Genesis Block  Nonce: ", block1.Nonce)
	fmt.Println("Genesis Block Difficulty: ", block1.Difficulty)
	fmt.Printf("Genesis Block Hash: %x\n", block1.Hash)
	fmt.Println("Subsequence Block  Timestamp: ", block2.Timestamp)
	fmt.Println("Subsequence  Block Nonce: ", block2.Nonce)
	fmt.Println("Subseuence Block  Difficulty: ", block2.Difficulty)
	fmt.Printf("Subsequence  Block  Hash: %x\n", block2.Hash)
	fmt.Printf("Subsequence  Block  Previous Hash(Genesis Block Hash): %x\n", block2.PrevHash)

	data, timestamp, nonce, difficulty := FetchDataFromBlocks(block1, block2)
	newBlock := NewBlock(data, block2.Hash, nonce, difficulty, timestamp)

	fmt.Println("Data of New Block: ", string(newBlock.Data))
	fmt.Println("Timestamp of New Block: ", newBlock.Timestamp)
	fmt.Println("Nonce of New Block: ", newBlock.Nonce)
	fmt.Println("Difficulty of New Block: ", newBlock.Difficulty)
	fmt.Printf("New Block Hash: %x\n", newBlock.Hash)
}
