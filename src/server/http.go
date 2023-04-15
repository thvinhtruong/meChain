package server

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func StartHTTPServer() {
	c := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	socket := GetSocket()

	socket.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Printf("connected: %v", s.ID())
		return nil
	})

	socket.OnEvent("/", "connect", func(s socketio.Conn) error {
		s.SetContext("")
		log.Printf("connected: %v", s.ID())
		return nil
	})

	router := gin.Default()
	router.Use(c)

	router.GET("/socket.io/", gin.WrapH(socket))
	router.POST("/socket.io/", gin.WrapH(socket))

	router.Run(":8000")
}
