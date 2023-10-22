package game

import (
	"fmt"

	"github.com/bamdadam/gotranj/internal/chess"
	"github.com/bamdadam/gotranj/internal/chess/piece"
	"github.com/bamdadam/gotranj/internal/chess/tile"
	"github.com/bamdadam/gotranj/internal/player"
)

type Board [][8]*tile.Tile

type Game struct {
	player1, player2         player.Player
	Board                    Board
	BlackPieces, WhitePieces []piece.Piece
}

func NewGame(player1, player2 player.Player) (*Game, error) {
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
	bp, err := chess.FillTilesWithPieces(true, gb)
	if err != nil {
		return nil, err
	}
	wp, err := chess.FillTilesWithPieces(false, gb)
	if err != nil {
		return nil, err
	}
	return &Game{
		player1:     player1,
		player2:     player2,
		Board:       gb,
		BlackPieces: bp,
		WhitePieces: wp,
	}, nil
}

func (b Board) String() string {
	res := ""
	for i := 0; i < 8; i++ {
		res += fmt.Sprintf("%d ", i)
		for j := 0; j < 8; j++ {
			res += fmt.Sprintf("%s ", b[i][j])
		}
		res += "\n"
	}
	res += "  a b c d e f g h"
	return res
}

// -1 means the piece can't move to the selected point
// (either because its out of the board or an ally piece resides there)
// 0 means the piece can move to the selected point
// 1 means the piece can move to the selected point and it will capture an enemy piece
func (b Board) CanMove(x, y int, isBlack bool) int {
	if x < 0 || x > 7 || y < 0 || y > 7 {
		return -1
	}
	if b[x][y].IsOccupied && b[x][y].Piece.IsBlack() == isBlack {
		return -1
	}
	if b[x][y].IsOccupied && b[x][y].Piece.IsBlack() != isBlack {
		return 1
	}
	return 0
}

func (g *Game) IsCheck(isBlack bool) bool {
	enemyPieces := g.WhitePieces
	if !isBlack {
		enemyPieces = g.BlackPieces
	}
	for _, p := range enemyPieces {
		for _, m := range p.GetMoves(p.GetX(), p.GetY(), g.Board.CanMove) {
			fmt.Println(p.String(), p.IsBlack(), m, fmt.Sprintf("%T", g.Board[m.X][m.Y].Piece))
			if g.Board[m.X][m.Y].IsOccupied {
				if _, ok := g.Board[m.X][m.Y].Piece.(*piece.King); ok && (g.Board[m.X][m.Y].Piece.IsBlack() == isBlack) {
					return true
				}
			}
		}
	}
	return false
}
