package entity

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"meChain2.0/src/blockchain/merkeltree"
	"meChain2.0/src/blockchain/utils"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := merkeltree.NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// CalculateHash calculates and returns block hash
func (b *Block) CalculateHash() []byte {
	blockData := bytes.Join(
		[][]byte{
			b.PrevBlockHash,
			b.HashTransactions(),
			utils.IntToHex(b.Timestamp),
			utils.IntToHex(int64(b.Nonce)),
			utils.IntToHex(int64(b.Height)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(blockData)

	return hash[:]
}

// IsValid checks whether block is valid
func (b *Block) IsValid(prevBlock *Block) bool {
	if bytes.Equal(b.PrevBlockHash, prevBlock.Hash) {
		return false
	}
	if bytes.Equal(b.Hash, b.CalculateHash()) {
		return false
	}
	return true
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func (b *Block) PrintBlock() {
	fmt.Printf("Block: %x\n", b.Hash)
	fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
	fmt.Printf("Height: %d\n", b.Height)
	fmt.Printf("Timestamp: %d\n", b.Timestamp)
	fmt.Printf("Nonce: %d\n", b.Nonce)
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
