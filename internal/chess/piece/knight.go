package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Knight struct {
	moveSet [][2]int
	isBlack bool
}

func (k *Knight) GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move {
	res := []move.Move{}
	for _, v := range k.moveSet {
		if canMove(x+v[0], y+v[1], k.isBlack) {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
		}
	}
	return res
}

func (k *Knight) String() string {
	if k.isBlack {
		return "♞"
	}
	return "♘"
}

func newKnight(isBlack bool) *Knight {
	ms := [][2]int{}
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if i != j {
				ms = append(ms, [2]int{i, j})
				ms = append(ms, [2]int{-i, j})
				ms = append(ms, [2]int{i, -j})
				ms = append(ms, [2]int{-i, -j})
			}
		}
	}
	return &Knight{
		moveSet: ms,
		isBlack: isBlack,
	}
}
