package game

import (
	"fmt"

	clib "github.com/MrNeocore/go-game-of-life/internal/cells"
)

func Run(rules clib.Rules, cellCount int, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan clib.Cell)
	cells := clib.MakeCells(cellCount)

	fmt.Printf("Starting cells: %v\n", cells)

	for i := 0; i < cellCount; i++ {
		go clib.RunCell(rules, cells[i], startChan, &cells, resultsChan)
	}

	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\tStep %d: ", i)
		cells = clib.NextStep(cellCount, startChan, &cells, resultsChan)
		fmt.Println(cells)
	}
}
