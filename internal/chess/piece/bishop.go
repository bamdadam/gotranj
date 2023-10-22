package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Bishop struct {
	moveSet map[[2]int]int
	isBlack bool
	x, y    int
}

func (b *Bishop) GetMoves(x, y int, canMove func(x, y int, isBlack bool) int) []move.Move {
	res := []move.Move{}
	for mv, iter := range b.moveSet {
		for i := 1; i < iter; i++ {
			mvRes := canMove(x+mv[0]*i, y+mv[1]*i, b.isBlack)
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
	ms := map[[2]int]int{}
	ms[[2]int{1, -1}] = 8
	ms[[2]int{-1, -1}] = 8
	ms[[2]int{1, 1}] = 8
	ms[[2]int{-1, 1}] = 8
	return &Bishop{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
