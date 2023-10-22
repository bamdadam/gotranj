package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Bishop struct {
	moveSet [][2]int
	isBlack bool
	x, y    int
}

func (b *Bishop) GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move {
	res := []move.Move{}
	for _, v := range b.moveSet {
		if canMove(x+v[0], y+v[1], b.isBlack) {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
		}
	}
	return res
}

func (b *Bishop) String() string {
	if !b.isBlack {
		return "♝"
	}
	return "♗"
}

func (b *Bishop) IsBlack() bool {
	return b.isBlack
}

func (b *Bishop) GetX() int {
	return b.x
}

func (b *Bishop) GetY() int {
	return b.y
}

func newBishop(isBlack bool, x, y int) *Bishop {
	ms := [][2]int{}
	for i := 1; i < 8; i++ {
		ms = append(ms, [2]int{i, i})
		ms = append(ms, [2]int{-i, -i})
		ms = append(ms, [2]int{-i, i})
		ms = append(ms, [2]int{i, -i})
	}
	return &Bishop{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
