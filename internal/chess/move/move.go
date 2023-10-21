package move

type Move struct {
	x, y int
}

func NewMove(x, y int) *Move {
	return &Move{
		x: x,
		y: y,
	}
}

func (m *Move) Convert(x, y int) int {
	return x*8 + y
}
