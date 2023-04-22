package server

// Share events for blockchain
type ShareEvent interface {
	GetName() string
	GetMessage() string
	GetRoom() int32
}

type IBlockChainEvent interface {
	GetPayload() []byte
}

type NewTransactionEvent struct {
	ShareEvent
	NodeID  int32
	Payload []byte // payload stores the transaction data
}

func (e *NewTransactionEvent) GetName() string {
	return "NewTransactionEvent"
}

func (e *NewTransactionEvent) GetMessage() string {
	return "New transaction received"
}

func (e *NewTransactionEvent) GetRoom() int32 {
	return e.NodeID
}

// GetPayload returns the serialized transaction data
func (e *NewTransactionEvent) GetPayload() []byte {
	return e.Payload
}

type NewBlockEvent struct {
	ShareEvent
	ProjectID int32
	Payload   []byte // payload stores miner transaction data

}

func (e *NewBlockEvent) GetName() string {
	return "NewBlockEvent"
}

func (e *NewBlockEvent) GetMessage() string {
	return "New block mined"
}

func (e *NewBlockEvent) GetRoom() int32 {
	return e.ProjectID
}

// GetPayload returns the serialized miner transaction data
func (e *NewBlockEvent) GetPayload() []byte {
	return e.Payload
}

type NewNodeEvent struct {
	ShareEvent
	Payload []byte // payload stores the node data
}

func (e *NewNodeEvent) GetName() string {
	return "NewNodeEvent"
}

func (e *NewNodeEvent) GetMessage() string {
	return "New node joined"
}

func (e *NewNodeEvent) GetRoom() int32 {
	return 0 // all subscribers
}

// GetPayload returns the serialized node data but node will not use this method
func (e *NewNodeEvent) GetPayload() []byte {
	return e.Payload
}
