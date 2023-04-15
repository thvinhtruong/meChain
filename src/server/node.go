package server

import (
	"time"

	"github.com/google/uuid"
	"meChain2.0/src/blockchain/entity"
)

type INode interface {
	GetID() uuid.UUID
	GetTransaction() []byte
	GetProjectID() int32
	GetPayload() []byte
	SetPayload(payload []byte)
	CreateNode(projectID int32, payload []byte, chain *entity.Blockchain)
	GetNode() Node
}

type Node struct {
	ID          uuid.UUID
	Transaction []byte // it is the serialized transactions
	ProjectID   int32
	Payload     []byte // payload stores serialized whole blockchain data
	Timestamp   int64
}

func NewNode(id uuid.UUID, projectID int32, payload []byte) *Node {
	return &Node{
		ID:          id,
		Transaction: nil,
		ProjectID:   projectID,
		Payload:     payload,
		Timestamp:   time.Now().Unix(),
	}
}

func (n *Node) CreateNode(projectID int32, payload []byte, chain *entity.Blockchain) {
	n.ID = uuid.New()
	n.ProjectID = projectID
	n.Payload = payload
	n.Timestamp = time.Now().Unix()
}

func (n *Node) GetNode() Node {
	return *n
}

func (n *Node) GetID() uuid.UUID {
	return n.ID
}

func (n *Node) GetTransaction() []byte {
	return n.Transaction
}

func (n *Node) GetProjectID() int32 {
	return n.ProjectID
}

// GetPayload returns the serialized blockchain data
func (n *Node) GetPayload() []byte {
	return n.Payload
}

func (n *Node) SetPayload(payload []byte) {
	n.Payload = payload
}
