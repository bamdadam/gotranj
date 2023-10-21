package game

import (
	"fmt"

	"github.com/bamdadam/gotranj/internal/chess"
	"github.com/bamdadam/gotranj/internal/chess/tile"
	"github.com/bamdadam/gotranj/internal/player"
)

type Board [][8]*tile.Tile

type Game struct {
	player1, player2 player.Player
	Board            Board
}

func NewGame(player1, player2 player.Player) *Game {
	gb := [][8]*tile.Tile{}
	for i := 0; i < 8; i++ {
		gb = append(gb, [8]*tile.Tile{})
		for j := 0; j < 8; j++ {
			gb[i][j] = &tile.Tile{
				IsOccupied: false,
				Piece:      nil,
			}
		}
	}
	chess.FillTilesWithPieces(true, gb)
	chess.FillTilesWithPieces(false, gb)
	return &Game{
		player1: player1,
		player2: player2,
		Board:   gb,
	}
}

func (b Board) String() string {
	res := ""
	for i := 0; i < 8; i++ {
		res += fmt.Sprintf("%d ", 8-i)
		for j := 0; j < 8; j++ {
			res += fmt.Sprintf("%s ", b[i][j])
		}
		res += "\n"
	}
	res += "  a b c d e f g h"
	return res
}
