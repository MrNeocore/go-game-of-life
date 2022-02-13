package cell

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MrNeocore/go-game-of-life/rules"
	"github.com/MrNeocore/go-game-of-life/state"
)

type Cell struct {
	Id    int
	State state.State
}

func (cell Cell) RunCell(cells *Cells1D) {
	for {
		<-cells.startChan

		neighbors := cell.getNeighbors(cells)
		numNeighborsAlive := countAliveCells(&neighbors)
		state := cell.getNewState(cells.Rules, numNeighborsAlive)

		cells.resultsChan <- Cell{Id: cell.Id, State: state}
	}
}

func (cell Cell) getNewState(rules rules.Rules, numNeighbors int) state.State {
	for _, neighbors := range rules[cell.State] {
		if numNeighbors == neighbors {
			return state.Alive
		}
	}

	return state.Dead
}

func (cell Cell) getNeighbors(cells *Cells1D) []Cell {
	leftId := cell.Id - 1

	if leftId == -1 {
		leftId = len(*cells.Cells) - 1
	}

	rightId := cell.Id + 1

	if rightId == len(*cells.Cells) {
		rightId = 0
	}

	leftCell := Cell{Id: leftId, State: (*cells.Cells)[leftId].State}
	rightCell := Cell{Id: rightId, State: (*cells.Cells)[rightId].State}

	return []Cell{leftCell, rightCell}
}

type Cells1D struct {
	Rules       rules.Rules
	CellCount   int
	Cells       *[]Cell
	startChan   chan bool
	resultsChan chan Cell
}

func NewCells(rules rules.Rules, cellCount int, startChan chan bool, resultsChan chan Cell) Cells1D {
	cells := Cells1D{
		Rules:       rules,
		CellCount:   cellCount,
		Cells:       makeCells(cellCount),
		startChan:   startChan,
		resultsChan: resultsChan,
	}
	cells.randomizeCells()

	return cells
}

func makeCells(cellCount int) *[]Cell {
	cells := make([]Cell, cellCount)

	return &cells
}

func (cells Cells1D) randomizeCells() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < cells.CellCount; i++ {
		(*cells.Cells)[i] = Cell{Id: i, State: state.State(rand.Intn(2) == 1)}
	}
}

func (cells Cells1D) String() string {
	out := ""

	for i := 0; i < cells.CellCount; i++ {
		cellState := (*cells.Cells)[i].State
		if cellState == state.Alive {
			out += "O "
		} else {
			out += "X "
		}
	}

	out += "\n"

	return out
}

func (cells *Cells1D) Start() {
	fmt.Println("=== Step 0 ===")
	fmt.Print(cells)

	for i := 0; i < cells.CellCount; i++ {
		go (*cells.Cells)[i].RunCell(cells)
	}
}

func (cells *Cells1D) Run(stepCount int) {
	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\n=== Step %d ===\n", i)
		cells.nextStep()
		fmt.Print(cells)
	}
}

func (cells *Cells1D) nextStep() {
	for i := 0; i < cells.CellCount; i++ {
		cells.startChan <- true
	}

	newCells := makeCells(cells.CellCount)

	for i := 0; i < cells.CellCount; i++ {
		newCell := <-cells.resultsChan
		(*newCells)[newCell.Id] = newCell
	}

	// Assign only after every Cell / goroutine has finished reading current states & produced its new state
	cells.Cells = newCells
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
