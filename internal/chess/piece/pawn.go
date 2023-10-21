package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Pawn struct {
	moveSet  [][2]int
	hasMoved bool
	isBlack  bool
}

func (p *Pawn) GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move {
	res := []move.Move{}
	for _, v := range p.moveSet {
		if canMove(x+v[0], y+v[1], p.isBlack) {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
		}
	}
	return res
}

func (p *Pawn) String() string {
	if p.isBlack {
		return "♟"
	}
	return "♙"
}

func newPawn(isBlack bool) *Pawn {
	ms := [][2]int{}
	if isBlack {
		ms = [][2]int{
			{1, 0}, {2, 0},
		}
	} else {
		ms = [][2]int{
			{-1, 0}, {-2, 0},
		}
	}
	return &Pawn{
		moveSet:  ms,
		hasMoved: false,
		isBlack:  isBlack,
	}
}
