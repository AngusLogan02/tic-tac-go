package ws

import (
	"log"
	"tic-tac-go/game"

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
				io.BroadcastToRoom("/stranger", roomName, "gameFound", roomList[roomMap[s.ID()]].player1)
			}
		}
		if roomFound == false {
			s.Join("game" + s.ID())
			roomMap[s.ID()] = len(roomList)
			roomList = append(roomList, room{
				roomID:    "game" + s.ID(),
				player1:   s.ID(),
				gamestate: game.InitialiseGamestate(),
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

func HandleMove(io *socketio.Server) {
	io.OnEvent("/stranger", "move", func(s socketio.Conn, location string) {
		currGame := roomList[roomMap[s.ID()]]
		var player string
		if s.ID() == currGame.player1 {
			player = "X"
		} else {
			player = "O"
		}

		if game.ValidMove(currGame.gamestate, location, player) {
			io.BroadcastToRoom("/stranger", currGame.roomID, "valid", location+player)
			winner := game.CheckWin(currGame.gamestate)
			if winner == "X" {
				io.BroadcastToRoom("/stranger", currGame.roomID, "winner", currGame.player1)
				io.ClearRoom("/stranger", currGame.roomID)
			} else if winner == "O" {
				io.BroadcastToRoom("/stranger", currGame.roomID, "winner", currGame.player2)
				io.ClearRoom("/stranger", currGame.roomID)
			}
		}
	})
}
