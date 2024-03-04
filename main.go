package main

// main file that contains all functions

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"/blockchain"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain
type Blockchain []Block

// NewBlock creates a new block
func NewBlock(index int, timestamp, data, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash calculates the hash of the block
func (b Block) calculateHash() string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash)))
	return hex.EncodeToString(hash.Sum(nil))
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := (*bc)[len(*bc)-1]
	newBlock := NewBlock(prevBlock.Index+1, time.Now().String(), data, prevBlock.Hash)
	*bc = append(*bc, newBlock)
}

// ModifyBlock modifies the data of a block
func (bc *Blockchain) ModifyBlock(index int, newData string) {
	if index >= 0 && index < len(*bc) {
		(*bc)[index].Data = newData
		(*bc)[index].Hash = (*bc)[index].calculateHash()
	}
}

// PrintBlockchain prints all blocks in the blockchain
func (bc Blockchain) PrintBlockchain() {
	for _, block := range bc {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("--------------")
	}
}

func main() {
	// Create the genesis block
	genesisBlock := NewBlock(0, time.Now().String(), "Genesis Block", "")

	// Create a new blockchain with the genesis block
	blockchain := Blockchain{genesisBlock}

	// Add some blocks to the blockchain
	blockchain.AddBlock("Block 1 Data")
	blockchain.AddBlock("Block 2 Data")

	// Print all blocks in the blockchain
	blockchain.PrintBlockchain()

	// Modify a block
	blockchain.ModifyBlock(1, "Modified Block 1 Data")

	// Print all blocks in the blockchain after modification
	fmt.Println("\nAfter modifying Block 1:")
	blockchain.PrintBlockchain()
}
