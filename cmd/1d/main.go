package main

import (
	"flag"
	"fmt"

	game1d "github.com/MrNeocore/go-game-of-life/game/1d"
	"github.com/MrNeocore/go-game-of-life/rules"
	"github.com/MrNeocore/go-game-of-life/state"
)

func parseCli() (cellCount int, stepCount int) {
	cellCountPtr := flag.Int("cellCount", 10, "Number of simulated cells")
	stepCountPtr := flag.Int("stepCount", 3, "Number of simulation steps")

	flag.Parse()

	return *cellCountPtr, *stepCountPtr
}

func main() {
	cellCount, stepCount := parseCli()

	fmt.Printf("Game of Life\n\n")

	var rules = rules.Rules{
		state.Alive: {1},
		state.Dead:  {1},
	}

	game1d.Run(rules, cellCount, stepCount)

	fmt.Printf("\nDone\n")
}
