package cell

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MrNeocore/go-game-of-life/dims"
	"github.com/MrNeocore/go-game-of-life/rules"
	"github.com/MrNeocore/go-game-of-life/state"
)

type Loc struct {
	X int
	Y int
}

type Cell struct {
	Loc   Loc
	State state.State
}

func (cell Cell) RunCell(cells Cells2D) {
	for {
		<-cells.startChan

		neighbors := cell.getNeighbors(cells)
		numNeighborsAlive := countAliveCells(&neighbors)
		state := cell.getNewState(cells.Rules, numNeighborsAlive)

		cells.resultsChan <- Cell{Loc: cell.Loc, State: state}
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

func (cell Cell) getNeighbors(cells Cells2D) []Cell {
	locs := []Loc{
		{X: cell.Loc.X - 1, Y: cell.Loc.Y - 1}, // topLeft
		{X: cell.Loc.X, Y: cell.Loc.Y - 1},     // top
		{X: cell.Loc.X + 1, Y: cell.Loc.Y - 1}, // topRight
		{X: cell.Loc.X + 1, Y: cell.Loc.Y},     // right
		{X: cell.Loc.X + 1, Y: cell.Loc.Y + 1}, // bottomRight
		{X: cell.Loc.X, Y: cell.Loc.Y + 1},     // bottom
		{X: cell.Loc.X - 1, Y: cell.Loc.Y + 1}, // bottomLeft
		{X: cell.Loc.X - 1, Y: cell.Loc.Y},     // left
	}

	neighbors := make([]Cell, 0, 8)

	for _, loc := range locs {
		if loc.X >= 0 && loc.X < cells.Dims.X && loc.Y >= 0 && loc.Y < cells.Dims.Y {
			neighbors = append(neighbors, (*cells.Cells)[loc.X][loc.Y])
		}
	}

	return neighbors
}

type Cells2D struct {
	Rules       rules.Rules
	Dims        dims.Dims
	Cells       *[][]Cell
	startChan   chan bool
	resultsChan chan Cell
}

func NewCells(rules rules.Rules, dims dims.Dims, startChan chan bool, resultsChan chan Cell) Cells2D {
	cells := Cells2D{
		Rules:       rules,
		Dims:        dims,
		Cells:       makeCells(dims),
		startChan:   startChan,
		resultsChan: resultsChan,
	}
	cells.randomizeCells()

	return cells
}

func makeCells(dims dims.Dims) *[][]Cell {
	_cells := make([][]Cell, dims.X)

	for i := range _cells {
		_cells[i] = make([]Cell, dims.Y)
	}

	return &_cells
}

func (cells Cells2D) randomizeCells() {
	rand.Seed(time.Now().UTC().UnixNano())

	for x := 0; x < cells.Dims.X; x++ {
		for y := 0; y < cells.Dims.Y; y++ {
			(*cells.Cells)[x][y] = Cell{Loc: Loc{X: x, Y: y}, State: state.State(rand.Intn(2) == 1)}
		}
	}
}

func (cells Cells2D) String() string {
	out := ""

	for x := 0; x < cells.Dims.X; x++ {
		for y := 0; y < cells.Dims.Y; y++ {
			cellState := (*cells.Cells)[x][y].State
			if cellState == state.Alive {
				out += "O "
			} else {
				out += "X "
			}
		}
		out += "\n"
	}

	return out
}

func (cells Cells2D) Start() {
	fmt.Println("=== Step 0 ===")
	fmt.Print(cells)

	for x := 0; x < cells.Dims.X; x++ {
		for y := 0; y < cells.Dims.Y; y++ {
			go (*cells.Cells)[x][y].RunCell(cells)
		}
	}
}

func (cells Cells2D) Run(stepCount int) {
	for i := 1; i < stepCount+1; i++ {
		fmt.Printf("\n=== Step %d ===\n", i)
		cells = cells.nextStep()
		fmt.Print(cells)
	}
}

func (cells Cells2D) nextStep() Cells2D {
	for i := 0; i < cells.Dims.X*cells.Dims.Y; i++ {
		cells.startChan <- true
	}

	newCells := Cells2D{
		Rules:       cells.Rules,
		Dims:        cells.Dims,
		Cells:       makeCells(cells.Dims),
		startChan:   cells.startChan,
		resultsChan: cells.resultsChan,
	}

	for i := 0; i < cells.Dims.X*cells.Dims.Y; i++ {
		newCell := <-cells.resultsChan
		(*newCells.Cells)[newCell.Loc.X][newCell.Loc.Y] = newCell
	}

	return newCells
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
