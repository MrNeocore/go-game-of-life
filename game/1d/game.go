package game

import (
	c1d "github.com/MrNeocore/go-game-of-life/internal/cell/1d"
	"github.com/MrNeocore/go-game-of-life/rules"
)

func Run(rules rules.Rules, cellCount int, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan c1d.Cell)

	cells := c1d.NewCells(rules, cellCount, startChan, resultsChan)

	cells.Start()
	cells.Run(stepCount)
}
