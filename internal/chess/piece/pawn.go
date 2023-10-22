package piece

import "github.com/bamdadam/gotranj/internal/chess/move"

type Pawn struct {
	moveSet  [][2]int
	hasMoved bool
	isBlack  bool
	x, y     int
}

func (p *Pawn) GetMoves(x, y int, canMove func(x, y int, isBlack bool) int) []move.Move {
	res := []move.Move{}
	for _, v := range p.moveSet {
		if canMove(x+v[0], y+v[1], p.isBlack) == 0 || canMove(x+v[0], y+v[1], p.isBlack) == 1 {
			mv := move.NewMove(x+v[0], y+v[1])
			res = append(res, *mv)
		}
	}
	return res
}

func (p *Pawn) String() string {
	if !p.isBlack {
		return "♟"
	}
	return "♙"
}

func (p *Pawn) IsBlack() bool {
	return p.isBlack
}

func (p *Pawn) GetX() int {
	return p.x
}

func (p *Pawn) GetY() int {
	return p.y
}

func newPawn(isBlack bool, x, y int) *Pawn {
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
		x:        x,
		y:        y,
	}
}
