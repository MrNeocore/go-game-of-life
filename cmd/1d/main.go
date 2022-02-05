package main

import (
	"fmt"

	game "github.com/MrNeocore/go-game-of-life/game/1d"
	"github.com/MrNeocore/go-game-of-life/rules"
	"github.com/MrNeocore/go-game-of-life/state"
)

const CELL_COUNT = 10
const STEP_COUNT = 2

func main() {
	fmt.Println("Game of Life")

	var rules = rules.Rules{
		state.Alive: {1},
		state.Dead:  {1},
	}

	game.Run(rules, CELL_COUNT, STEP_COUNT)

	fmt.Println("Done")
}
