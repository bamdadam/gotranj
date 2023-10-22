package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Queen struct {
	moveSet map[[2]int]int
	isBlack bool
	x, y    int
}

func (q *Queen) GetMoves(x, y int, canMove func(x, y int, isBlack bool) int) []move.Move {
	res := []move.Move{}
	for mv, iter := range q.moveSet {
		for i := 1; i < iter; i++ {
			mvRes := canMove(x+mv[0]*i, y+mv[1]*i, q.isBlack)
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

func (q *Queen) String() string {
	if !q.isBlack {
		return "♛"
	}
	return "♕"
}

func (q *Queen) IsBlack() bool {
	return q.isBlack
}

func (q *Queen) GetX() int {
	return q.x
}

func (q *Queen) GetY() int {
	return q.y
}

func newQueen(isBlack bool, x, y int) *Queen {
	ms := map[[2]int]int{}
	ms[[2]int{1, 0}] = 8
	ms[[2]int{-1, 0}] = 8
	ms[[2]int{0, 1}] = 8
	ms[[2]int{0, -1}] = 8
	ms[[2]int{1, -1}] = 8
	ms[[2]int{-1, -1}] = 8
	ms[[2]int{1, 1}] = 8
	ms[[2]int{-1, 1}] = 8

	return &Queen{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
