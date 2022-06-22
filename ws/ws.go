package ws

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func HandleOnConnect(io *socketio.Server) {
	io.OnConnect("/stranger", func(s socketio.Conn) error {
		s.Leave(s.ID())

		roomFound := false
		for _, room := range io.Rooms("/stranger") {
			if io.RoomLen("/stranger", room) == 1 {
				s.Join(room)
				roomFound = true
			}
		}
		if roomFound == false {
			s.Join("game" + s.ID())
		}
		log.Println(io.Rooms("/stranger"))
		return nil
	})
}

func HandleOnDisconnect(io *socketio.Server) {
	io.OnDisconnect("/", func(s socketio.Conn, reason string) {})
}

func HandleOnDisconnectStranger(io *socketio.Server) {
	io.OnDisconnect("/stranger", func(s socketio.Conn, reason string) {
		log.Println(s.ID(), "disconnected from", s.Namespace(), ":", reason)
	})
}

func HandleMsgEvent(io *socketio.Server) {
	io.OnEvent("/stranger", "msg", func(s socketio.Conn, msg string) string {
		log.Println("message:", msg)
		s.SetContext(msg)
		return msg
	})
}
