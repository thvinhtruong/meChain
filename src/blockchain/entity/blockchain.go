package entity

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

var (
	// list of valid miners
	miners = []string{"vincent", "kd", "ton", "ntl", "duyen"}
)

type Blockchain struct {
	Blocks              []*Block
	PendingTransactions []*Transaction
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock(&Transaction{
		FromAddress: []byte{},
		ToAddress:   []byte{},
		Amount:      0,
		Timestamp:   time.Now().Unix(),
	})

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) GetTopBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddPendingTransaction(transaction *Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
}

// should be fixed
func (bc *Blockchain) MineBlock(minerAddress string) *Block {
	// check if miner is valid
	isValidMiner := false
	for _, miner := range miners {
		if miner == minerAddress {
			isValidMiner = true
			break
		}
	}
	if !isValidMiner {
		log.Panic("Invalid miner")
	}

	tx := &Transaction{
		ID:          []byte(uuid.New().String()),
		FromAddress: []byte(minerAddress),
		ToAddress:   []byte("zeus"),
		Amount:      0,
	}

	// we simply get the top of the block to get the prehash
	lasthash := bc.GetTopBlock().Hash
	lastheight := bc.GetTopBlock().Height
	bc.PendingTransactions = append(bc.PendingTransactions, tx)

	// Create new block
	lastheight++
	newBlock := NewBlock(bc.PendingTransactions, lasthash, lastheight)
	bc.AddBlock(newBlock)

	return newBlock
}

func (bc *Blockchain) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(bc)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlockchain(data []byte) *Blockchain {
	var blockchain Blockchain

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&blockchain)
	if err != nil {
		log.Panic(err)
	}

	return &blockchain
}

func (bc *Blockchain) PrintChain() {
	for i, block := range bc.Blocks {
		fmt.Printf("============ Block %d ============\n", i)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
	}
}

func (bc *Blockchain) PrintAllTransactions() {
	for _, block := range bc.Blocks {
		for i, tx := range block.Transactions {
			fmt.Printf("============ Transaction %d ============\n", i)
			fmt.Printf("ID: %x\n", tx.ID)
			fmt.Printf("From: %s\n", tx.FromAddress)
			fmt.Printf("To: %s\n", tx.ToAddress)
			fmt.Printf("Amount: %d\n", tx.Amount)
			fmt.Printf("Timestamp: %d\n", tx.Timestamp)
		}
	}

}
