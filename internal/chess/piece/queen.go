package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Queen struct {
	moveSet [][2]int
	isBlack bool
	x, y    int
}

func (q *Queen) GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move {
	res := []move.Move{}
	for _, v := range q.moveSet {
		if canMove(x+v[0], y+v[1], q.isBlack) {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
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
	ms := [][2]int{}
	for i := 1; i < 8; i++ {
		ms = append(ms, [2]int{i, i})
		ms = append(ms, [2]int{-i, -i})
		ms = append(ms, [2]int{-i, i})
		ms = append(ms, [2]int{i, -i})
		ms = append(ms, [2]int{i, 0})
		ms = append(ms, [2]int{-i, 0})
		ms = append(ms, [2]int{0, i})
		ms = append(ms, [2]int{0, -i})
	}
	return &Queen{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
