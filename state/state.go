package state

type State bool

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
