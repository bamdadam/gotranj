package move

type Move struct {
	X, Y int
}

func NewMove(x, y int) *Move {
	return &Move{
		X: x,
		Y: y,
	}
}

func (m *Move) Convert(x, y int) int {
	return x*8 + y
}
