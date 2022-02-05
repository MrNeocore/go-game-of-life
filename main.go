package main

import (
	"fmt"
	"math/rand"
	"time"
)

type State bool

const (
	Alive State = true
	Dead  State = false
)

type Cell struct {
	id    int
	state State
}

func (s State) String() string {
	if s == Alive {
		return "Alive"
	} else {
		return "Dead"
	}
}

var rules = map[State][]int{
	Alive: {1},
	Dead:  {1},
}

func newState(currentState State, numNeighbours int) State {
	for _, neighbours := range rules[currentState] {
		if numNeighbours == neighbours {
			return Alive
		}
	}

	return Dead
}

func makeCells(count int) []Cell {
	cells := make([]Cell, count)

	for i := range cells {
		cells[i] = Cell{id: i, state: State(rand.Intn(2) == 1)}
	}

	return cells
}

func countAliveCells(cells []Cell) int {
	aliveCount := 0

	for _, cell := range cells {
		if cell.state == Alive {
			aliveCount += 1
		}
	}

	return aliveCount
}

func getNeighbours(self Cell, cells *[]Cell) []Cell {
	leftId := self.id - 1

	if leftId == -1 {
		leftId = len(*cells) - 1
	}

	rightId := self.id + 1

	if rightId == len(*cells) {
		rightId = 0
	}

	leftCell := Cell{id: leftId, state: (*cells)[leftId].state}
	rightCell := Cell{id: rightId, state: (*cells)[rightId].state}

	return []Cell{leftCell, rightCell}
}

func runCell(self Cell, startChan chan bool, cells *[]Cell, resultsChan chan Cell) {
	for {
		<-startChan

		neighbours := getNeighbours(self, cells)
		numNeighbours := countAliveCells(neighbours)
		state := newState(self.state, numNeighbours)

		resultsChan <- Cell{id: self.id, state: state}
	}
}

func nextStep(startChan chan bool, cells *[]Cell, resultsChan chan Cell) []Cell {
	for i := 0; i < CELL_COUNT; i++ {
		startChan <- true
	}

	newCells := make([]Cell, CELL_COUNT)

	for i := 0; i < CELL_COUNT; i++ {
		newCell := <-resultsChan
		newCells[newCell.id] = newCell
	}

	return newCells
}

const CELL_COUNT = 10
const STEPS_COUNT = 2

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Game of Life")

	cells := makeCells(CELL_COUNT)
	fmt.Printf("Starting cells: %v\n", cells)

	startChan := make(chan bool)
	resultsChan := make(chan Cell)

	for i := 0; i < CELL_COUNT; i++ {
		go runCell(cells[i], startChan, &cells, resultsChan)
	}

	for i := 1; i < STEPS_COUNT+1; i++ {
		fmt.Printf("\tStep %d: ", i)
		cells = nextStep(startChan, &cells, resultsChan)
		fmt.Println(cells)
	}

	fmt.Println("Done")
}
