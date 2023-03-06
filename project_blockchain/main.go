package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	Hash         []byte
	PrevHash     []byte
	Nonce        int
	Difficulty   int
	Transactions []Transaction
}

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp int64
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	nonce := []byte(strconv.Itoa(b.Nonce))
	difficulty := []byte(strconv.Itoa(b.Difficulty))
	transactions, _ := json.Marshal(b.Transactions)
	blockData := append(b.Data, b.PrevHash...)
	blockData = append(blockData, timestamp...)
	blockData = append(blockData, nonce...)
	blockData = append(blockData, difficulty...)
	blockData = append(blockData, transactions...)
	hash := sha256.Sum256(blockData)
	b.Hash = hash[:]
}

func (b *Block) GetData() []byte {
	return b.Data
}

func (b *Block) GetTransactions() []Transaction {
	return b.Transactions
}

func NewBlock(data []byte, prevHash []byte, nonce int, difficulty int, timestamp int64, transactions []Transaction) *Block {
	block := &Block{timestamp, data, []byte{}, prevHash, nonce, difficulty, transactions}
	block.SetHash()
	return block
}

func FetchDataFromBlocks(block1, block2 *Block) ([]byte, int64, int, int, []Transaction) {
	data := append(block1.Data, block2.Data...)
	timestamp := time.Now().Unix()
	nonce := block1.Nonce + block2.Nonce
	difficulty := block1.Difficulty + block2.Difficulty
	transactions := make([]Transaction, 0)
	transactions = append(transactions, block1.Transactions...)
	transactions = append(transactions, block2.Transactions...)
	return data, timestamp, nonce, difficulty, transactions
}

func main() {
	transaction1 := Transaction{
		Sender:    "Sender name is :Hemel Hasan",
		Receiver:  " Receiver name is:Niloy Islam",
		Amount:    100,
		Timestamp: time.Now().Unix(),
	}
	transaction2 := Transaction{
		Sender:    " Sender name is:Niloy Islam",
		Receiver:  " Receiver Name is:Hemel Hasan",
		Amount:    50,
		Timestamp: time.Now().Unix(),
	}
	block1 := NewBlock([]byte("Hemel Hasan Transafer 100$ to Niloy's Wallet and "), []byte{}, 12, 3, time.Now().Unix(), []Transaction{transaction1})
	block2 := NewBlock([]byte("Niloy Islam Has Received 100$ from Hemel Hasan"), block1.Hash, 15, 4, time.Now().Unix(), []Transaction{transaction2})
	fmt.Println("Genesis Block  Timestamp: ", block1.Timestamp)
	fmt.Println("Genesis Block  Nonce: ", block1.Nonce)
	fmt.Println("Genesis Block Difficulty: ", block1.Difficulty)
	fmt.Printf("Genesis Block Hash: %x\n", (block1.Hash))
	fmt.Printf("Genesis Block Hash: %x\n", len(block1.Hash))
	fmt.Println("Transactions of Genesis Block: ", block1.Transactions)
	fmt.Println("Subsequence Block  Timestamp: ", block2.Timestamp)
	fmt.Println("Subsequence  Block Nonce: ", block2.Nonce)
	fmt.Println("Subsequence Block  Difficulty: ", block2.Difficulty)
	fmt.Printf("Subsequence  Block  Hash: %x\n", block2.Hash)
	fmt.Printf("Subsequence  Block  Previous Hash(Genesis Block Hash): %x\n", block2.PrevHash)
	fmt.Println("Transactions of Subsequence Block: ", block2.Transactions)

	data, timestamp, nonce, difficulty, transactions := FetchDataFromBlocks(block1, block2)
	newBlock := NewBlock(data, block2.Hash, nonce, difficulty, timestamp, transactions)

	fmt.Println("Data of New Block: ", string(newBlock.Data))
	fmt.Println("Timestamp of New Block: ", newBlock.Timestamp)
	fmt.Println("Nonce of New Block: ", newBlock.Nonce)
	fmt.Println("Difficulty of New Block: ", newBlock.Difficulty)
	fmt.Println("Transactions of New Block: ", newBlock.Transactions)
	fmt.Printf("New Block Hash: %x\n", newBlock.Hash)
	fmt.Printf("Previous Hash(Subsequence Block Hash): %x\n", newBlock.PrevHash)
}
