package game

import (
	clib "github.com/MrNeocore/go-game-of-life/2d/internal/cell"
	"github.com/MrNeocore/go-game-of-life/dims"
	"github.com/MrNeocore/go-game-of-life/rules"
)

func Run(rules rules.Rules, dims dims.Dims, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan clib.Cell)

	cells := clib.NewCells(rules, dims, startChan, resultsChan)

	cells.Start()
	cells.Run(stepCount)
}
