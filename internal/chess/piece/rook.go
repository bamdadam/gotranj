package piece

import (
	"github.com/bamdadam/gotranj/internal/chess/move"
)

type Rook struct {
	moveSet map[[2]int]int
	isBlack bool
	x, y    int
}

func (r *Rook) GetMoves(x, y int, canMove func(x, y int, isBlack bool) int) []move.Move {
	res := []move.Move{}
	for mv, iter := range r.moveSet {
		for i := 1; i < iter; i++ {
			mvRes := canMove(x+mv[0]*i, y+mv[1]*i, r.isBlack)
			if mvRes == 0 {
				mv := move.NewMove(x+mv[0]*i, y+mv[1]*i)
				res = append(res, *mv)
			} else if mvRes == 1 {
				mv := move.NewMove(x+mv[0]*i, y+mv[1]*i)
				res = append(res, *mv)
				break
			} else {
				break
			}
		}
	}
	return res
}

func (r *Rook) String() string {
	if !r.isBlack {
		return "♜"
	}
	return "♖"
}

func (r *Rook) IsBlack() bool {
	return r.isBlack
}

func (r *Rook) GetX() int {
	return r.x
}

func (r *Rook) GetY() int {
	return r.y
}

func newRook(isBlack bool, x, y int) *Rook {
	ms := map[[2]int]int{}
	ms[[2]int{1, 0}] = 8
	ms[[2]int{-1, 0}] = 8
	ms[[2]int{0, 1}] = 8
	ms[[2]int{0, -1}] = 8

	return &Rook{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
