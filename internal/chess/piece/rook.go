package piece

import (
	"github.com/bamdadam/gotranj/internal/chess/move"
)

type Rook struct {
	moveSet [][2]int
	isBlack bool
}

func (r *Rook) GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move {
	res := []move.Move{}
	for _, v := range r.moveSet {
		if canMove(x+v[0], y+v[1], r.isBlack) {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
		}
	}
	return res
}

func (r *Rook) String() string {
	if r.isBlack {
		return "♜"
	}
	return "♖"
}

func newRook(isBlack bool) *Rook {
	ms := [][2]int{}
	for i := 1; i < 8; i++ {
		ms = append(ms, [2]int{i, 0})
		ms = append(ms, [2]int{-i, 0})
		ms = append(ms, [2]int{0, i})
		ms = append(ms, [2]int{0, -i})
	}
	return &Rook{
		moveSet: ms,
		isBlack: isBlack,
	}
}
