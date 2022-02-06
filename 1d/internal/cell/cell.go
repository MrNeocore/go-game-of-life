package cell

import (
	"math/rand"
	"time"

	"github.com/MrNeocore/go-game-of-life/rules"
	"github.com/MrNeocore/go-game-of-life/state"
)

// CELL
type Cell struct {
	Id    int
	State state.State
}

func (cell Cell) getNewState(rules rules.Rules, numNeighbors int) state.State {
	for _, neighbors := range rules[cell.State] {
		if numNeighbors == neighbors {
			return state.Alive
		}
	}

	return state.Dead
}

func (cell Cell) getNeighbors(cells *[]Cell) []Cell {
	leftId := cell.Id - 1

	if leftId == -1 {
		leftId = len(*cells) - 1
	}

	rightId := cell.Id + 1

	if rightId == len(*cells) {
		rightId = 0
	}

	leftCell := Cell{Id: leftId, State: (*cells)[leftId].State}
	rightCell := Cell{Id: rightId, State: (*cells)[rightId].State}

	return []Cell{leftCell, rightCell}
}

func (cell Cell) RunCell(rules rules.Rules, startChan chan bool, cells *[]Cell, resultsChan chan Cell) {
	for {
		<-startChan

		neighbors := cell.getNeighbors(cells)
		numneighbors := countAliveCells(&neighbors)
		state := cell.getNewState(rules, numneighbors)

		resultsChan <- Cell{Id: cell.Id, State: state}
	}
}

func MakeCells(count int) []Cell {
	rand.Seed(time.Now().UTC().UnixNano())

	cells := make([]Cell, count)

	for i := range cells {
		cells[i] = Cell{Id: i, State: state.State(rand.Intn(2) == 1)}
	}

	return cells
}

func countAliveCells(cells *[]Cell) int {
	aliveCount := 0

	for _, cell := range *cells {
		if cell.State == state.Alive {
			aliveCount += 1
		}
	}

	return aliveCount
}

func NextStep(startChan chan bool, cells *[]Cell, resultsChan chan Cell) *[]Cell {
	for i := 0; i < len(*cells); i++ {
		startChan <- true
	}

	newCells := make([]Cell, len(*cells))

	for i := 0; i < len(*cells); i++ {
		newCell := <-resultsChan
		newCells[newCell.Id] = newCell
	}

	return &newCells
}
