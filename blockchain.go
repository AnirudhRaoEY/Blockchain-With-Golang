package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Defining a Block- Main Structure
type Block struct {
	nounce       int
	previousHash string
	timestamp    int64
	transation   []string
}

// Creating a new Block
func NewBlock(nonce int, previousHash string) *Block {
	//We are Intializing a New Block to variable b
	b := new(Block)
	//We are Initializing a New Block using the paremeters of the function
	b.timestamp = time.Now().UnixNano()
	b.nounce = nonce
	b.previousHash = previousHash
	return b
}
func (b *Block) Print() {
	fmt.Printf("timestamp   %d\n", b.timestamp)
	fmt.Printf("Nounce   %d\n", b.nounce)
	fmt.Printf("PreviousHash   %s\n", b.previousHash)
	fmt.Printf("Transaction   %s\n", b.transation)
}

// Blockchain Struct- To Provide Basic Structure
type Blockchain struct {
	transationPool []string
	chain          []*Block
}

// Creating a New Blockchain function
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

// Creating a Chain of Blocks=- i.e Blockchain function
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// Current Blockchain returns a sequence
// We use Print Function to return the basic Blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("=", 25))

}

// Logging Function
func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	//Creating a New Block
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.Print()
	blockchain.CreateBlock(2, "hash 2")
	blockchain.Print()
}
