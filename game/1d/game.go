package game

import (
	"fmt"

	clib "github.com/MrNeocore/go-game-of-life/internal/cells"
)

func startCells(rules clib.Rules, startChan chan bool, cells *[]clib.Cell, resultsChan chan clib.Cell) {
	fmt.Printf("Starting cells: %v\n", cells)

	for i := 0; i < len(*cells); i++ {
		go clib.RunCell(rules, (*cells)[i], startChan, cells, resultsChan)
	}
}

func runCellSteps(stepCount int, startChan chan bool, cells *[]clib.Cell, resultsChan chan clib.Cell) {
	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\tStep %d: ", i)
		cells = clib.NextStep(len(*cells), startChan, cells, resultsChan)
		fmt.Println(cells)
	}
}

func Run(rules clib.Rules, cellCount int, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan clib.Cell)

	cells := clib.MakeCells(cellCount)

	startCells(rules, startChan, &cells, resultsChan)

	runCellSteps(stepCount, startChan, &cells, resultsChan)
}
