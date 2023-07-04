package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Block Struct
type Block struct {
	Index int
	Timestamp int64
	Data string
	PrevHash string
	Hash string
	Nonce int
	Difficulty int
	Transactions []*Transaction
}

// Transaction Struct
type Transaction struct {
	sender string
	receiver string
	amount int
}

// Blockchain Struct
type Blockchain struct {
	blocks []*Block
}

// Create a new block
func NewBlock(index int, timestamp int64, data string, prevHash string, difficulty int) *Block {
	block := &Block {
		Index: index,
		Timestamp: timestamp,
		Data: data,
		PrevHash: prevHash,
		Difficulty: difficulty,
	}
	block.Mine()
	return block
}

// Calculate the hash value of the block
func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10) + b.Data + b.PrevHash + strconv.Itoa(b.Nonce) + b.TransactionHash()
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
} 

// Mine provide the proof-of-work
func (b *Block) Mine() {
	target := bytes.Repeat([]byte("0"), b.Difficulty)
	for {
		b.Hash = b.CalculateHash()
		if bytes.HasPrefix([]byte(b.Hash), target) {
			break
		} 
		b.Nonce++
	}
}

// Add a new block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, time.Now().Unix(), data, prevBlock.Hash, prevBlock.Difficulty)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create first block of blockchain
func CreateGenesisBlock() *Block {
	return NewBlock(0, time.Now().Unix(), "Genesis Block", "", 2)
}

// Create a new instance of a blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{CreateGenesisBlock()},
	}
}

func(bc *Blockchain) IsChainValid() bool {
	for i:=1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		prevBlock: bc.blocks[i-1]
	}
}

func main() {
	
}