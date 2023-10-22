package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type King struct {
	moveSet map[[2]int]int
	isBlack bool
	x, y    int
}

func (k *King) GetMoves(x, y int, canMove func(x, y int, isBlack bool) int) []move.Move {
	res := []move.Move{}
	for mv, iter := range k.moveSet {
		for i := 1; i < iter; i++ {
			mvRes := canMove(x+mv[0]*i, y+mv[1]*i, k.isBlack)
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

func (k *King) String() string {
	if !k.isBlack {
		return "♚"
	}
	return "♔"
}

func (k *King) IsBlack() bool {
	return k.isBlack
}

func (k *King) GetX() int {
	return k.x
}

func (k *King) GetY() int {
	return k.y
}

func newKing(isBlack bool, x, y int) *King {
	ms := map[[2]int]int{}
	ms[[2]int{1, 0}] = 2
	ms[[2]int{-1, 0}] = 2
	ms[[2]int{0, 1}] = 2
	ms[[2]int{0, -1}] = 2
	ms[[2]int{1, -1}] = 2
	ms[[2]int{-1, -1}] = 2
	ms[[2]int{1, 1}] = 2
	ms[[2]int{-1, 1}] = 2
	return &King{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
