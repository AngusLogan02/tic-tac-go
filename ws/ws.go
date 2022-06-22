package ws

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func HandleOnConnect(io *socketio.Server) {
	io.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("session id:", s.ID())
		log.Println("current client count:", io.Count())

		return nil
	})
}

func HandleOnDisconnect(io *socketio.Server) {
	io.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnect:", reason)
		log.Println("current client count:", io.Count())
	})
}

func HandleMsgEvent(io *socketio.Server) {
	io.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		log.Println("message:", msg)
		s.SetContext(msg)
		return msg
	})
}
