package server

type Networks struct {
	pubsub  ISocketPubSub
	node    INode
	handler IEventHandler
}

func NewNetworks(pubsub ISocketPubSub, node INode, handler IEventHandler) *Networks {
	return &Networks{
		pubsub:  pubsub,
		node:    node,
		handler: handler,
	}
}

func (n *Networks) EventListener() {
	n.pubsub.Subscribe("NewTransactionEvent", n.handler.HandleNewTransaction)
	n.pubsub.Subscribe("NewBlockEvent", n.handler.HandleNewBlock)
	n.pubsub.Subscribe("NewNodeEvent", n.handler.HandleNewNode)
}

func (n *Networks) EventPublisher() {
	n.pubsub.Publish("NewTransactionEvent", n.node)
	n.pubsub.Publish("NewBlockEvent", n.node)
	n.pubsub.Publish("NewNodeEvent", n.node)
}

func (n *Networks) Start() {
	n.EventListener()
	n.EventPublisher()
}
