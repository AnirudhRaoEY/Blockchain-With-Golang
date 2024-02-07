package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const Minning_difficulty = 3

// Defining a Block- Main Structure
type Block struct {
	timestamp    int64
	nounce       int
	previousHash [32]byte

	transactions []*Transaction
}

// Creating a new Block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	//We are Intializing a New Block to variable b
	b := new(Block)
	//We are Initializing a New Block using the paremeters of the function
	b.timestamp = time.Now().UnixNano()
	b.nounce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}
func (b *Block) Print() {
	fmt.Printf("timestamp   %d\n", b.timestamp)
	fmt.Printf("Nounce   %d\n", b.nounce)
	fmt.Printf("PreviousHash   %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

// Creating a Hash: Using the Sha256 Algorithm
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// Json Marshalling
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nounce       int            `json:"nounce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transaction  []*Transaction `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nounce:       b.nounce,
		PreviousHash: b.previousHash,
		Transaction:  b.transactions,
	})

}

// Blockchain Struct- To Provide Basic Structure
type Blockchain struct {
	transationPool []*Transaction
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
	b := NewBlock(nonce, previousHash, bc.transationPool)
	bc.chain = append(bc.chain, b)
	bc.transationPool = []*Transaction{}
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
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transationPool = append(bc.transationPool, t)

}

func (bc *Blockchain) CopytranscationPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transationPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockChainAddress,
				t.recipientBlockChainAddress,
				t.value))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nounce int, previousHash [32]byte, trnsactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessblock := Block{0, nounce, previousHash, trnsactions}
	guessHashstr := fmt.Sprintf("%x", guessblock.Hash())
	return guessHashstr[:difficulty] == zeros
}

// Proof of Work
func (bc *Blockchain) ProofofWork() int {
	transactions := bc.CopytranscationPool()
	previousHash := bc.LastBlock().Hash()
	nounce := 0
	//Commputing the correct Nounce
	for bc.ValidProof(nounce, previousHash, transactions, Minning_difficulty) {
		nounce += 1
	}
	return nounce
}

// Creating a Transaction Block
type Transaction struct {
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}

// Creating a new Transcation Block
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Printing the Transaction
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" Sender_Blockchain_Address %s\n", t.senderBlockChainAddress)
	fmt.Printf(" Recipient_Blockchain_Address %s\n", t.recipientBlockChainAddress)
	fmt.Printf(" Value of Transaction         %.1f\n", t.value)
}

// Marshalling the Json
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockChainAddress    string  `json:"senderBlockChainAddress"`
		RecipientBlockChainAddress string  `json:"recipientBlockChainAddress"`
		Value                      float32 `json:"value"`
	}{
		//Passing Values
		SenderBlockChainAddress:    t.senderBlockChainAddress,
		RecipientBlockChainAddress: t.recipientBlockChainAddress,
		Value:                      t.value,
	})
}

// Logging Function
func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	//Creating a New Block
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().previousHash
	nounce := blockchain.ProofofWork()
	blockchain.CreateBlock(nounce, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("B", "C", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	previousHash = blockchain.LastBlock().Hash()
	nounce = blockchain.ProofofWork()
	blockchain.CreateBlock(nounce, previousHash)
	blockchain.Print()

}
