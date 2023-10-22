package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type King struct {
	moveSet [][2]int
	isBlack bool
	x, y    int
}

func (k *King) GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move {
	res := []move.Move{}
	for _, v := range k.moveSet {
		if canMove(x+v[0], y+v[1], k.isBlack) {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
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
	ms := [][2]int{}
	// generate all possible moves for a king
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// exclude the case where both i and j are 0
			if i != 0 || j != 0 {
				ms = append(ms, [2]int{i, j})
			}
		}
	}
	return &King{
		moveSet: ms,
		isBlack: isBlack,
		x:       x,
		y:       y,
	}
}
