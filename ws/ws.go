package ws

import (
	"log"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
)

func HandleOnConnect(io *socketio.Server) {
	io.OnConnect("/stranger", func(s socketio.Conn) error {
		s.LeaveAll()
		s.Join("stranger_waiting")
		log.Println(strconv.Itoa(io.RoomLen("/stranger", "stranger_waiting")) + " strangers waiting")
		return nil
	})
}

func HandleOnDisconnect(io *socketio.Server) {
	io.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnect:", reason)
	})
}

func HandleMsgEvent(io *socketio.Server) {
	io.OnEvent("/stranger", "msg", func(s socketio.Conn, msg string) string {
		log.Println("message:", msg)
		s.SetContext(msg)
		return msg
	})
}
