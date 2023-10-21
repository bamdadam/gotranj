package tile

import (
	"fmt"

	"github.com/bamdadam/gotranj/internal/chess/piece"
)

type Tile struct {
	IsOccupied bool
	Piece      piece.Piece
}

func (t Tile) String() string {
	if t.IsOccupied {
		return fmt.Sprintf("%s", t.Piece)
	}
	// return empty square character as a string
	return "□"
}
