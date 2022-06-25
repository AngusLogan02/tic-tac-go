package ws

import (
	"strconv"
	"tic-tac-go/game"

	socketio "github.com/googollee/go-socket.io"
)

type room struct {
	roomID    string
	namespace string
	player1   string
	player2   string
	gamestate [][]string
	movecount int
}

var roomMap = make(map[string]int)
var roomList []room

func HandleStrangerConnect(io *socketio.Server) {
	io.OnConnect("/stranger", func(s socketio.Conn) error {
		s.Leave(s.ID())

		roomFound := false
		for _, roomName := range io.Rooms("/stranger") {
			if io.RoomLen("/stranger", roomName) == 1 {
				s.Join(roomName)
				for i := range roomList {
					if roomList[i].roomID == roomName {
						roomMap[s.ID()] = i
						roomList[i].player2 = s.ID()
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
				namespace: "/stranger",
				player1:   s.ID(),
				gamestate: game.InitialiseGamestate(),
				movecount: 0,
			})
		}
		return nil
	})
}

func HandleFriendConnect(io *socketio.Server) {
	io.OnConnect("/friend", func(s socketio.Conn) error {
		// s.Leave(s.ID())
		return nil
	})
}

func OnReceiveFriendID(io *socketio.Server) {
	io.OnEvent("/friend", "friendID", func(s socketio.Conn, friendID string) {
		if io.RoomLen("/friend", friendID) == 0 {
			s.Join(friendID)
			roomMap[s.ID()] = len(roomList)
			roomList = append(roomList, room{
				roomID:    friendID,
				namespace: "/friend",
				player1:   s.ID(),
				gamestate: game.InitialiseGamestate(),
				movecount: 0,
			})
			return
		} else if io.RoomLen("/friend", friendID) == 1 {
			s.Join(friendID)
			for i := range roomList {
				if roomList[i].roomID == friendID {
					roomMap[s.ID()] = i
					roomList[i].player2 = s.ID()
					break
				}
			}
			io.BroadcastToRoom("/friend", friendID, "gameFound", roomList[roomMap[s.ID()]].player1)
			return
		} else {
			io.BroadcastToNamespace("/", "gameFull", s.ID())
			delete(roomMap, s.ID())
		}
	})
}

func HandleDisconnect(io *socketio.Server) {
	io.OnDisconnect("/", func(s socketio.Conn, reason string) {
		currGame := roomList[roomMap[s.ID()]]
		io.BroadcastToRoom(currGame.namespace, currGame.roomID, "dc")
		io.ClearRoom(currGame.namespace, currGame.roomID)
	})
}

func HandleMove(io *socketio.Server) {
	io.OnEvent("/", "move", func(s socketio.Conn, location string) {
		currGame := roomList[roomMap[s.ID()]]
		var player string
		if s.ID() == currGame.player1 {
			player = "X"
		} else {
			player = "O"
		}

		x, _ := strconv.Atoi(location[0:1])
		y, _ := strconv.Atoi(location[1:2])

		valid, winner := game.Move(currGame.gamestate, x, y, player)

		if valid {
			io.BroadcastToRoom(currGame.namespace, currGame.roomID, "valid", location+player)
			roomList[roomMap[s.ID()]].movecount++
			if roomList[roomMap[s.ID()]].movecount == 9 && winner == "" {
				io.BroadcastToRoom(currGame.namespace, currGame.roomID, "draw")
			}
			if winner == "X" {
				io.BroadcastToRoom(currGame.namespace, currGame.roomID, "winner", currGame.player1)
				io.ClearRoom(currGame.namespace, currGame.roomID)
			} else if winner == "O" {
				io.BroadcastToRoom(currGame.namespace, currGame.roomID, "winner", currGame.player2)
				io.ClearRoom(currGame.namespace, currGame.roomID)
			}
		}
	})
}
