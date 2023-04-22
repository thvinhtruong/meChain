package server

import (
	"bytes"
	"log"

	"meChain2.0/src/blockchain/entity"
)

// Events between blockchain and clients include:
/// 1. New transaction
/// 2. New block (miner client)
/// 3. New node (self-defined term)

type IEventHandler interface {
	HandleNewTransaction(node INode)
	HandleNewBlock(node INode)
	HandleNewNode(node INode)
}

type EventHandler struct {
	event IBlockChainEvent
}

func NewEventHandler(event IBlockChainEvent) *EventHandler {
	return &EventHandler{
		event: event,
	}
}

func (e *EventHandler) HandleNewTransaction(node INode) {
	// deserialize payload data
	incomingTransaction := e.event.GetPayload() // transaction
	data := entity.DeserializeTransactions(incomingTransaction)
	blockchain := entity.DeserializeBlockchain(node.GetPayload())

	// check if transaction is valid
	if bytes.Equal(incomingTransaction, node.GetTransaction()) {
		log.Println("Invalid transaction happends at handler: ", node.GetID())
		return
	}

	// add transaction to the pending transaction pool
	for _, transaction := range data {
		blockchain.AddPendingTransaction(&transaction)
	}

}

func (e *EventHandler) HandleNewBlock(node INode) {
	// deserialize payload data
	incomingMiner := e.event.GetPayload() // miner

	// add block to the blockchain
	blockchain := entity.DeserializeBlockchain(node.GetPayload())
	if blockchain == nil {
		log.Println("Invalid blockchain happends at handler: ", node.GetID())
		return
	}

	newblock := blockchain.MineBlock(string(incomingMiner))
	if newblock == nil {
		log.Println("Invalid block happends at handler: ", node.GetID())
		return
	}

	node.SetPayload(blockchain.Serialize())

}

func (e *EventHandler) HandleNewNode(node INode) {
	// deserialize payload data
	sentBlockchain := e.event.GetPayload() // blockchain data

	// create new node
	node.CreateNode(node.GetProjectID(), sentBlockchain, nil)

	// receive blockchain from master
	blockchain := entity.DeserializeBlockchain(sentBlockchain)
	if blockchain == nil {
		log.Println("Invalid blockchain happends at handler: ", node.GetID())
		return
	}
	// save node to db
	// node.SaveNode()

}
