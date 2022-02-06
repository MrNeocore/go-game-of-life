package rules

import "github.com/MrNeocore/go-game-of-life/state"

// Mapping of cell State to number of neighbors that will lead to State = Alive
type Rules map[state.State][]int
