package entity

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"math/big"
	"time"

	"github.com/gofrs/uuid"
)

type Transaction struct {
	ID          []byte
	FromAddress []byte
	ToAddress   []byte
	Amount      int
	Timestamp   int64
}

// should fix this
func NewTransaction(fromAddress, toAddress []byte, amount int) *Transaction {
	return &Transaction{
		uuid.Must(uuid.NewV4()).Bytes(),
		fromAddress,
		toAddress,
		amount,
		time.Now().Unix(),
	}
}

// Serialize returns a serialized Transaction
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

// SetID sets ID of a Transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TrimmedCopy creates a trimmed copy of Transaction to be used in signing
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []byte

	inputs = append(inputs, tx.FromAddress...)

	var outputs []byte
	outputs = append(outputs, tx.ToAddress...)

	txCopy := Transaction{tx.ID, inputs, outputs, tx.Amount, tx.Timestamp}

	return txCopy
}

// Sign signs each input of a Transaction
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey) ([]byte, error) {
	txCopy := tx.TrimmedCopy()

	hashedTx := txCopy.Hash()
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, hashedTx)
	if err != nil {
		return nil, err
	}

	signature := append(r.Bytes(), s.Bytes()...)

	return signature, nil
}

// Verify verifies signature of Transaction input
func (tx *Transaction) Verify(pubKey ecdsa.PublicKey, signature []byte) bool {
	txCopy := tx.TrimmedCopy()

	hashedTx := txCopy.Hash()
	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])

	ok := ecdsa.Verify(&pubKey, hashedTx, &r, &s)
	return ok
}

// DeserializeTransaction deserializes a transaction
func DeserializeTransactions(data []byte) []Transaction {
	var transactions []Transaction

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transactions)
	if err != nil {
		log.Panic(err)
	}

	return transactions
}
