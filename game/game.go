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

func CheckWin(gamestate [][]string) string {
	for i := range gamestate {
		if gamestate[0][i] == gamestate[1][i] && gamestate[1][i] == gamestate[2][i] {
			// row check
			return gamestate[0][i]
		}
		if gamestate[i][0] == gamestate[i][1] && gamestate[i][1] == gamestate[i][2] {
			// column check
			return gamestate[i][0]
		}
	}
	if gamestate[0][0] == gamestate[1][1] && gamestate[1][1] == gamestate[2][2] {
		// diagonal check (\)
		return gamestate[0][0]
	}
	if gamestate[0][2] == gamestate[1][1] && gamestate[1][1] == gamestate[2][0] {
		// diagonal check (/)
		return gamestate[0][2]
	}
	return ""
}
