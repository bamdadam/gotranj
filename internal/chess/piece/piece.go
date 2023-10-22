package piece

import (
	"errors"
	"fmt"

	"github.com/bamdadam/gotranj/internal/chess/move"
)

type Piece interface {
	GetMoves(x, y int, canMove func(x, y int, isBlack bool) int) []move.Move
	String() string
	IsBlack() bool
	GetX() int
	GetY() int
}

func NewPiece(pieceType string, x, y int) (Piece, error) {
	switch pieceType {
	case "bp":
		return newPawn(true, x, y), nil
	case "wp":
		return newPawn(false, x, y), nil
	case "br":
		return newRook(true, x, y), nil
	case "wr":
		return newRook(false, x, y), nil
	case "bb":
		return newBishop(true, x, y), nil
	case "wb":
		return newBishop(false, x, y), nil
	case "bq":
		return newQueen(true, x, y), nil
	case "wq":
		return newQueen(false, x, y), nil
	case "bk":
		return newKing(true, x, y), nil
	case "wk":
		return newKing(false, x, y), nil
	case "bn":
		return newKnight(true, x, y), nil
	case "wn":
		return newKnight(false, x, y), nil
	}
	return nil, errors.New(fmt.Sprintf("piece of type %s is not available", pieceType))
}
