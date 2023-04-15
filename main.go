package main

import (
	"github.com/google/uuid"
	"meChain2.0/src/blockchain/entity"
	"meChain2.0/src/server"
)

var (
	SocketServer = server.GetSocket()
)

func blockchain() *entity.Blockchain {
	return entity.NewBlockchain()
}

func main() {
	// init blockchain
	chain := blockchain()

	// init pubsub
	pubsub := server.NewSocketPubSub(SocketServer)

	// init start node
	startNode := server.NewNode(uuid.New(), 0, chain.Serialize())

	// init event handler
	eventHandler := server.NewEventHandler(startNode)

	// init network
	network := server.NewNetworks(pubsub, startNode, eventHandler)

	// start network
	network.Start()

	go SocketServer.Serve()
	defer SocketServer.Close()

	// start http server
	server.StartHTTPServer()

}
