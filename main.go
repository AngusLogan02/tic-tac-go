package main

import (
	"log"
	"tic-tac-go/handlers"
	"tic-tac-go/ws"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)

	io := socketio.NewServer(nil)

	ws.HandleFriendConnect(io)
	ws.OnReceiveFriendID(io)
	ws.HandleStrangerConnect(io)
	ws.HandleDisconnect(io)
	ws.HandleMove(io)

	r.LoadHTMLGlob("public/html/*")
	r.Static("/static/css", "./public/css")
	r.Static("/static/js", "./public/js")

	r.GET("/", handlers.IndexHandler)
	r.GET("/friend/:friendID", handlers.FriendHandler)
	r.GET("/stranger", handlers.StrangerHandler)

	r.GET("/socket.io/*any", gin.WrapH(io))
	r.POST("/socket.io/*any", gin.WrapH(io))

	go func() {
		if err := io.Serve(); err != nil {
			log.Fatal("socket.io listen server error:", err)
		}
	}()
	defer io.Close()

	r.Run()
}
