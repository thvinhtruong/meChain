package server

import (
	socketio "github.com/googollee/go-socket.io"
)

// Socket interface for pubsub events handling
type ISocketPubSub interface {
	Subscribe(topic string, handler func(node INode))
	Publish(topic string, node INode)
}

// SocketPubSub struct
type SocketPubSub struct {
	socket *socketio.Server
}

// NewSocketPubSub creates a new SocketPubSub instance
func NewSocketPubSub(socket *socketio.Server) *SocketPubSub {
	return &SocketPubSub{
		socket: socket,
	}
}

// Subscribe to a topic
func (s *SocketPubSub) Subscribe(event string, handler func(node INode)) {
	s.socket.OnEvent("/", event, func(s socketio.Conn, msg interface{}) {
		s.Emit(event, msg)
		handler(msg.(INode))
	})
}

// Publish to a topic
func (s *SocketPubSub) Publish(event string, node INode) {
	s.socket.BroadcastToNamespace("/", event, node.GetNode())
}
