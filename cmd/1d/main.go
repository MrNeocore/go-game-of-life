package main

import (
	"fmt"

	clib "github.com/MrNeocore/go-game-of-life/internal/cells"
)

const CELL_COUNT = 10
const STEP_COUNT = 2

func run(cellCount int, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan clib.Cell)
	cells := clib.MakeCells(cellCount)

	fmt.Printf("Starting cells: %v\n", cells)

	for i := 0; i < cellCount; i++ {
		go clib.RunCell(cells[i], startChan, &cells, resultsChan)
	}

	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\tStep %d: ", i)
		cells = clib.NextStep(CELL_COUNT, startChan, &cells, resultsChan)
		fmt.Println(cells)
	}
}

func main() {
	fmt.Println("Game of Life")

	run(CELL_COUNT, STEP_COUNT)

	fmt.Println("Done")
}
