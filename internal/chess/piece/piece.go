package piece

import (
	"errors"
	"fmt"

	"github.com/bamdadam/gotranj/internal/chess/move"
)

type Piece interface {
	GetMoves(x, y int, canMove func(x, y int, isBlack bool) bool) []move.Move
	String() string
}

func NewPiece(pieceType string) (Piece, error) {
	switch pieceType {
	case "bp":
		return newPawn(true), nil
	case "wp":
		return newPawn(false), nil
	case "br":
		return newRook(true), nil
	case "wr":
		return newRook(false), nil
	case "bb":
		return newBishop(true), nil
	case "wb":
		return newBishop(false), nil
	case "bq":
		return newQueen(true), nil
	case "wq":
		return newQueen(false), nil
	case "bk":
		return newKing(true), nil
	case "wk":
		return newKing(false), nil
	case "bn":
		return newKnight(true), nil
	case "wn":
		return newKnight(false), nil
	}
	return nil, errors.New(fmt.Sprintf("piece of type %s is not available", pieceType))
}
