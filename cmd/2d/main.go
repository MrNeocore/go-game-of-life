package main

import (
	"flag"
	"fmt"

	"github.com/MrNeocore/go-game-of-life/2d/game"
	_dims "github.com/MrNeocore/go-game-of-life/dims"
	"github.com/MrNeocore/go-game-of-life/rules"
	"github.com/MrNeocore/go-game-of-life/state"
)

func parseCli() (dims _dims.Dims, stepCount int) {
	cellCountXPtr := flag.Int("X", 10, "Number of simulated cells (X)")
	cellCountYPtr := flag.Int("Y", 10, "Number of simulated cells (Y)")
	stepCountPtr := flag.Int("steps", 3, "Number of simulation steps")

	flag.Parse()

	return _dims.Dims{X: *cellCountXPtr, Y: *cellCountYPtr}, *stepCountPtr
}

func main() {
	dims, stepCount := parseCli()

	fmt.Printf("Game of Life\n\n")

	var rules = rules.Rules{
		state.Alive: {2, 3},
		state.Dead:  {2},
	}

	game.Run(rules, dims, stepCount)

	fmt.Printf("\nDone\n")
}
