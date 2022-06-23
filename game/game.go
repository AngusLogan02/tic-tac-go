package game

import (
	"strconv"
)

func ValidMove(gamestate [][]string, location string, player string) bool {
	x, _ := strconv.Atoi(location[0:1])
	y, _ := strconv.Atoi(location[1:2])
	var valid bool

	if gamestate[x][y] == "" {
		gamestate[x][y] = player
		valid = true
	} else {
		valid = false
	}
	return valid
}

func InitialiseGamestate() [][]string {
	gamestate := make([][]string, 3)
	for i := range gamestate {
		gamestate[i] = make([]string, 3)
	}
	return gamestate
}
