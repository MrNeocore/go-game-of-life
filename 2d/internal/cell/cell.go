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
	X     int
	Y     int
	State state.State
}

// TODO: Use
// type Cells struct {
// 	dims  dims.Dims
// 	cells *[][]Cell
// }

func (cell Cell) getNewState(rules rules.Rules, numNeighbors int) state.State {
	for _, neighbors := range rules[cell.State] {
		if numNeighbors == neighbors {
			return state.Alive
		}
	}

	return state.Dead
}

func (cell Cell) getneighbors(cells *[][]Cell, dims dims.Dims) []Cell {
	topLeft := Loc{X: cell.X - 1, Y: cell.Y - 1}
	top := Loc{X: cell.X, Y: cell.Y - 1}
	topRight := Loc{X: cell.X + 1, Y: cell.Y - 1}
	right := Loc{X: cell.X + 1, Y: cell.Y}
	bottomRight := Loc{X: cell.X + 1, Y: cell.Y + 1}
	bottom := Loc{X: cell.X, Y: cell.Y + 1}
	bottomLeft := Loc{X: cell.X - 1, Y: cell.Y + 1}
	left := Loc{X: cell.X - 1, Y: cell.Y}

	locs := []Loc{
		topLeft,
		top,
		topRight,
		right,
		bottomRight,
		bottom,
		bottomLeft,
		left,
	}

	neighbors := make([]Cell, 0, 8)

	for _, loc := range locs {
		if loc.X >= 0 && loc.X < dims.X && loc.Y >= 0 && loc.Y < dims.Y {
			neighbors = append(neighbors, (*cells)[loc.X][loc.Y])
		}
	}

	return neighbors
}

func (cell Cell) RunCell(rules rules.Rules, startChan chan bool, cells *[][]Cell, resultsChan chan Cell, dims dims.Dims) {
	for {
		<-startChan

		neighbors := cell.getneighbors(cells, dims)
		numneighborsAlive := countAliveCells(&neighbors)
		state := cell.getNewState(rules, numneighborsAlive)

		resultsChan <- Cell{X: cell.X, Y: cell.Y, State: state}
	}
}

func InitCells(dims dims.Dims) [][]Cell {
	cells := makeCells(dims)
	randomizeCells(&cells, dims)

	return cells
}

func makeCells(dims dims.Dims) [][]Cell {
	cells := make([][]Cell, dims.Y)

	for i := range cells {
		cells[i] = make([]Cell, dims.X)
	}

	return cells
}

func randomizeCells(cells *[][]Cell, dims dims.Dims) {
	rand.Seed(time.Now().UTC().UnixNano())

	for x := 0; x < dims.X; x++ {
		for y := 0; y < dims.Y; y++ {
			(*cells)[x][y] = Cell{X: x, Y: y, State: state.State(rand.Intn(2) == 1)}
		}
	}
}

func PrintCells(cells *[][]Cell, dims dims.Dims) {
	for x := 0; x < dims.X; x++ {
		for y := 0; y < dims.Y; y++ {
			cellState := (*cells)[x][y].State
			if cellState == state.Alive {
				fmt.Print("O ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Println("")
	}
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

func NextStep(dims dims.Dims, startChan chan bool, cells *[][]Cell, resultsChan chan Cell) *[][]Cell {
	for i := 0; i < dims.X*dims.Y; i++ {
		startChan <- true
	}

	newCells := makeCells(dims)

	for i := 0; i < dims.X*dims.Y; i++ {
		newCell := <-resultsChan
		newCells[newCell.X][newCell.Y] = newCell
	}

	return &newCells
}
