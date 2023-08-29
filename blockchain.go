package main

import (
	"fmt"
	"log"
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

// Logging Function
func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	//Creating a New Block
	b := NewBlock(0, "init Hash")
	b.Print()
}
