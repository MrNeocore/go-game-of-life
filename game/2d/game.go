package game

import (
	"github.com/MrNeocore/go-game-of-life/dims"
	c2d "github.com/MrNeocore/go-game-of-life/internal/cell/2d"
	"github.com/MrNeocore/go-game-of-life/rules"
)

func Run(rules rules.Rules, dims dims.Dims, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan c2d.Cell)

	cells := c2d.NewCells(rules, dims, startChan, resultsChan)

	cells.Start()
	cells.Run(stepCount)
}
