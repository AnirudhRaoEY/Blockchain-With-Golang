package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Defining a Block- Main Structure
type Block struct {
	nounce       int
	previousHash [32]byte
	timestamp    int64
	transation   []string
}

// Creating a new Block
func NewBlock(nonce int, previousHash [32]byte) *Block {
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
	fmt.Printf("PreviousHash   %x\n", b.previousHash)
	fmt.Printf("Transaction   %s\n", b.transation)
}

// Creating a Hash: Using the Sha256 Algorithm
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

// Json Marshalling
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nounce       int      `json:"nounce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transaction  []string `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nounce:       b.nounce,
		PreviousHash: b.previousHash,
		Transaction:  b.transation,
	})

}

// Blockchain Struct- To Provide Basic Structure
type Blockchain struct {
	transationPool []string
	chain          []*Block
}

// Creating a New Blockchain function
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Creating a Chain of Blocks=- i.e Blockchain function
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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
//Creating a Transaction Block
type Transaction struct{
	senderBlockChainAddress  string
	recipientBlockChainAddress string
	value float32
}
//Creating a new Transcation Block
func NewTransaction(sender string, recipient string , value float32 ) *Transaction{
	return &Transaction{sender, recipient,value}
}
//Printing the Transaction

// Logging Function
func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	//Creating a New Block
	blockchain := NewBlockchain()
	blockchain.Print()
	previousHash := blockchain.LastBlock().previousHash
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()

}
