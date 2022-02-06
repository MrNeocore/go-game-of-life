package game

import (
	"fmt"

	clib "github.com/MrNeocore/go-game-of-life/2d/internal/cell"
	"github.com/MrNeocore/go-game-of-life/dims"
	"github.com/MrNeocore/go-game-of-life/rules"
)

func Run(rules rules.Rules, dims dims.Dims, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan clib.Cell)

	cells := clib.InitCells(dims)

	startCells(rules, dims, startChan, &cells, resultsChan)

	runCellSteps(stepCount, dims, startChan, &cells, resultsChan)
}

func startCells(rules rules.Rules, dims dims.Dims, startChan chan bool, cells *[][]clib.Cell, resultsChan chan clib.Cell) {
	fmt.Println("=== Step 0 ===")
	clib.PrintCells(cells, dims)

	for x := 0; x < dims.X; x++ {
		for y := 0; y < dims.Y; y++ {
			go (*cells)[x][y].RunCell(rules, startChan, cells, resultsChan, dims)
		}
	}
}

func runCellSteps(stepCount int, dims dims.Dims, startChan chan bool, cells *[][]clib.Cell, resultsChan chan clib.Cell) {
	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\n=== Step %d ===\n", i)
		cells = clib.NextStep(dims, startChan, cells, resultsChan)
		clib.PrintCells(cells, dims)
	}
}
