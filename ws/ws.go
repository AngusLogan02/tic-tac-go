package ws

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type room struct {
	roomID    string
	player1   string
	player2   string
	gamestate [][]string
}

var roomMap = make(map[string]int)
var roomList []room

func HandleOnConnect(io *socketio.Server) {
	io.OnConnect("/stranger", func(s socketio.Conn) error {
		s.Leave(s.ID())

		roomFound := false
		for _, roomName := range io.Rooms("/stranger") {
			if io.RoomLen("/stranger", roomName) == 1 {
				s.Join(roomName)
				for i := range roomList {
					if roomList[i].roomID == roomName {
						roomMap[s.ID()] = i
						break
					}
				}
				roomFound = true
				io.BroadcastToRoom("/stranger", roomName, "gameFound", "found")
			}
		}
		if roomFound == false {
			s.Join("game" + s.ID())
			roomMap[s.ID()] = len(roomList)
			roomList = append(roomList, room{
				roomID:  "game" + s.ID(),
				player1: s.ID(),
			})
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

func HandleTestEvent(io *socketio.Server) {
	io.OnEvent("/stranger", "testEvent", func(s socketio.Conn, sID string) {
		roomIndex := roomMap[sID]
		socketRoom := roomList[roomIndex]
		log.Println("socket id is", sID, "therefore they are in room", socketRoom.roomID)
	})
}
