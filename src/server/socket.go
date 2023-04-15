package server

import (
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

var (
	Instance *socketio.Server
	once     sync.Once
)

func init() {
	once.Do(func() {
		Instance = socketio.NewServer(nil)
	})
}

func GetSocket() *socketio.Server {
	return Instance
}
