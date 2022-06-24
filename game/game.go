package game

func Move(gamestate [][]string, x int, y int, player string) (bool, string) {
	var valid bool

	if gamestate[x][y] == "" {
		gamestate[x][y] = player
		valid = true
		for i := 0; i < 3; i++ {
			if gamestate[x][i] != player {
				break
			}
			if i == 2 {
				return valid, player
			}
		}

		for i := 0; i < 3; i++ {
			if gamestate[i][y] != player {
				break
			}
			if i == 2 {
				return valid, player
			}
		}

		if x == y {
			for i := 0; i < 3; i++ {
				if gamestate[i][i] != player {
					break
				}
				if i == 2 {
					return valid, player
				}
			}
		}

		if x+y == 2 {
			for i := 0; i < 3; i++ {
				if gamestate[i][2-i] != player {
					break
				}
				if i == 2 {
					return valid, player
				}
			}
		}
	} else {
		valid = false
	}
	return valid, ""
}

func InitialiseGamestate() [][]string {
	gamestate := make([][]string, 3)
	for i := range gamestate {
		gamestate[i] = make([]string, 3)
	}
	return gamestate
}
