package cells

import (
	"math/rand"
	"time"
)

type State bool

// Mapping of cell State to number of neighbours that will lead to State = Alive
type Rules map[State][]int

const (
	Alive State = true
	Dead  State = false
)

func (state State) String() string {
	if state == Alive {
		return "Alive"
	} else {
		return "Dead"
	}
}

type Cell struct {
	Id    int
	State State
}

func newState(rules Rules, currentState State, numNeighbours int) State {
	for _, neighbours := range rules[currentState] {
		if numNeighbours == neighbours {
			return Alive
		}
	}

	return Dead
}

func MakeCells(count int) []Cell {
	rand.Seed(time.Now().UTC().UnixNano())

	cells := make([]Cell, count)

	for i := range cells {
		cells[i] = Cell{Id: i, State: State(rand.Intn(2) == 1)}
	}

	return cells
}

func CountAliveCells(cells []Cell) int {
	aliveCount := 0

	for _, cell := range cells {
		if cell.State == Alive {
			aliveCount += 1
		}
	}

	return aliveCount
}

func getNeighbours(self Cell, cells *[]Cell) []Cell {
	leftId := self.Id - 1

	if leftId == -1 {
		leftId = len(*cells) - 1
	}

	rightId := self.Id + 1

	if rightId == len(*cells) {
		rightId = 0
	}

	leftCell := Cell{Id: leftId, State: (*cells)[leftId].State}
	rightCell := Cell{Id: rightId, State: (*cells)[rightId].State}

	return []Cell{leftCell, rightCell}
}

func RunCell(rules Rules, self Cell, startChan chan bool, cells *[]Cell, resultsChan chan Cell) {
	for {
		<-startChan

		neighbours := getNeighbours(self, cells)
		numNeighbours := CountAliveCells(neighbours)
		state := newState(rules, self.State, numNeighbours)

		resultsChan <- Cell{Id: self.Id, State: state}
	}
}

func NextStep(cellCount int, startChan chan bool, cells *[]Cell, resultsChan chan Cell) []Cell {
	for i := 0; i < cellCount; i++ {
		startChan <- true
	}

	newCells := make([]Cell, cellCount)

	for i := 0; i < cellCount; i++ {
		newCell := <-resultsChan
		newCells[newCell.Id] = newCell
	}

	return newCells
}
