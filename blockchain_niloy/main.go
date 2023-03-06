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
	ExtraData  map[string]string
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	nonce := []byte(strconv.Itoa(b.Nonce))
	difficulty := []byte(strconv.Itoa(b.Difficulty))
	extraData := []byte(fmt.Sprintf("%v", b.ExtraData))
	blockData := append(b.Data, b.PrevHash...)
	blockData = append(blockData, timestamp...)
	blockData = append(blockData, nonce...)
	blockData = append(blockData, difficulty...)
	blockData = append(blockData, extraData...)
	hash := sha256.Sum256(blockData)
	b.Hash = hash[:]
}

func (b *Block) GetData() []byte {
	return b.Data
}

func (b *Block) GetExtraData() map[string]string {
	return b.ExtraData
}

func NewBlock(data []byte, prevHash []byte, nonce int, difficulty int, timestamp int64, extraData map[string]string) *Block {
	block := &Block{timestamp, data, []byte{}, prevHash, nonce, difficulty, extraData}
	block.SetHash()
	return block
}

func FetchDataFromBlocks(block1, block2 *Block) ([]byte, int64, int, int, map[string]string) {
	data := append(block1.Data, block2.Data...)
	timestamp := time.Now().Unix()
	nonce := block1.Nonce + block2.Nonce
	difficulty := block1.Difficulty + block2.Difficulty
	extraData := make(map[string]string)
	for k, v := range block1.ExtraData {
		extraData[k] = v
	}
	for k, v := range block2.ExtraData {
		extraData[k] = v
	}
	return data, timestamp, nonce, difficulty, extraData
}

func main() {
	extraData1 := map[string]string{"TransactionID": "12345", "Sender": "Hemel Hasan", "Receiver": "Niloy Islam"}
	extraData2 := map[string]string{"TransactionID": "54321", "Sender": "Niloy Islam", "Receiver": "Hemel Hasan"}
	block1 := NewBlock([]byte("Hemel Hasan Transafer 100$ to Niloy's Wallet and "), []byte{}, 12, 3, time.Now().Unix(), extraData1)
	block2 := NewBlock([]byte("Niloy Islam Has Received 100$ from Hemel Hasan"), block1.Hash, 15, 4, time.Now().Unix(), extraData2)
	fmt.Println("Genesis Block  Timestamp: ", block1.Timestamp)
	fmt.Println("Genesis Block  Nonce: ", block1.Nonce)
	fmt.Println("Genesis Block Difficulty: ", block1.Difficulty)
	fmt.Printf("Genesis Block Hash: %x\n", block1.Hash)
	fmt.Println("ExtraData of Genesis Block: ", block1.ExtraData)
	fmt.Println("Subseuence Block  Timestamp: ", block2.Timestamp)
	fmt.Println("Subseuence  Block Nonce: ", block2.Nonce)
	fmt.Println("Subseuence Block  Difficulty: ", block2.Difficulty)
	fmt.Printf("Subseuence  Block  Hash: %x\n", block2.Hash)
	fmt.Println("ExtraData of Subseuence Block: ", block2.ExtraData)
	fmt.Printf("Subseuence  Block  Previous Hash(Genesis Block Hash): %x\n", block2.PrevHash)

	data, timestamp, nonce, difficulty, extraData := FetchDataFromBlocks(block1, block2)
	newBlock := NewBlock(data, block2.Hash, nonce, difficulty, timestamp, extraData)

	fmt.Println("Data of New Block: ", string(newBlock.Data))
	fmt.Println("Timestamp of New Block: ", newBlock.Timestamp)
	fmt.Println("Nonce of New Block: ", newBlock.Nonce)
	fmt.Println("Difficulty of New Block: ", newBlock.Difficulty)
	fmt.Println("ExtraData of New Block: ", newBlock.ExtraData)
	fmt.Printf("New Block Hash: %x\n", newBlock.Hash)
}
