package main

import (
	"fmt"

	clib "github.com/MrNeocore/go-game-of-life/internal/cells"
)

const CELL_COUNT = 10
const STEP_COUNT = 2

func run(rules clib.Rules, cellCount int, stepCount int) {
	startChan := make(chan bool)
	resultsChan := make(chan clib.Cell)
	cells := clib.MakeCells(cellCount)

	fmt.Printf("Starting cells: %v\n", cells)

	for i := 0; i < cellCount; i++ {
		go clib.RunCell(rules, cells[i], startChan, &cells, resultsChan)
	}

	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\tStep %d: ", i)
		cells = clib.NextStep(CELL_COUNT, startChan, &cells, resultsChan)
		fmt.Println(cells)
	}
}

func main() {
	fmt.Println("Game of Life")

	var rules = clib.Rules{
		clib.Alive: {1},
		clib.Dead:  {1},
	}

	run(rules, CELL_COUNT, STEP_COUNT)

	fmt.Println("Done")
}
