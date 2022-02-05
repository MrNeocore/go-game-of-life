package main

import (
	"fmt"

	game "github.com/MrNeocore/go-game-of-life/game/1d"
	clib "github.com/MrNeocore/go-game-of-life/internal/cells"
)

const CELL_COUNT = 10
const STEP_COUNT = 2

func main() {
	fmt.Println("Game of Life")

	var rules = clib.Rules{
		clib.Alive: {1},
		clib.Dead:  {1},
	}

	game.Run(rules, CELL_COUNT, STEP_COUNT)

	fmt.Println("Done")
}
