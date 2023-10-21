package chess

import (
	"github.com/bamdadam/gotranj/internal/chess/piece"
	"github.com/bamdadam/gotranj/internal/chess/tile"
)

func FillTilesWithPieces(isBlack bool, gameBoard [][8]*tile.Tile) error {
	if isBlack {
		p, err := piece.NewPiece("br")
		if err != nil {
			return err
		}
		gameBoard[0][0].Piece = p
		gameBoard[0][0].IsOccupied = true
		p, err = piece.NewPiece("bn")
		if err != nil {
			return err
		}
		gameBoard[0][1].Piece = p
		gameBoard[0][1].IsOccupied = true
		p, err = piece.NewPiece("bb")
		if err != nil {
			return err
		}
		gameBoard[0][2].Piece = p
		gameBoard[0][2].IsOccupied = true
		p, err = piece.NewPiece("bq")
		if err != nil {
			return err
		}
		gameBoard[0][3].Piece = p
		gameBoard[0][3].IsOccupied = true
		p, err = piece.NewPiece("bk")
		if err != nil {
			return err
		}
		gameBoard[0][4].Piece = p
		gameBoard[0][4].IsOccupied = true
		p, err = piece.NewPiece("bb")
		if err != nil {
			return err
		}
		gameBoard[0][5].Piece = p
		gameBoard[0][5].IsOccupied = true
		p, err = piece.NewPiece("bn")
		if err != nil {
			return err
		}
		gameBoard[0][6].Piece = p
		gameBoard[0][6].IsOccupied = true
		p, err = piece.NewPiece("br")
		if err != nil {
			return err
		}
		gameBoard[0][7].Piece = p
		gameBoard[0][7].IsOccupied = true
		p, err = piece.NewPiece("bp")
		if err != nil {
			return err
		}
		for i := 0; i < 8; i++ {
			gameBoard[1][i].Piece = p
			gameBoard[1][i].IsOccupied = true
		}
	} else {
		p, err := piece.NewPiece("wr")
		if err != nil {
			return err
		}
		gameBoard[7][0].Piece = p
		gameBoard[7][0].IsOccupied = true
		p, err = piece.NewPiece("wn")
		if err != nil {
			return err
		}
		gameBoard[7][1].Piece = p
		gameBoard[7][1].IsOccupied = true
		p, err = piece.NewPiece("wb")
		if err != nil {
			return err
		}
		gameBoard[7][2].Piece = p
		gameBoard[7][2].IsOccupied = true
		p, err = piece.NewPiece("wq")
		if err != nil {
			return err
		}
		gameBoard[7][3].Piece = p
		gameBoard[7][3].IsOccupied = true
		p, err = piece.NewPiece("wk")
		if err != nil {
			return err
		}
		gameBoard[7][4].Piece = p
		gameBoard[7][4].IsOccupied = true
		p, err = piece.NewPiece("wb")
		if err != nil {
			return err
		}
		gameBoard[7][5].Piece = p
		gameBoard[7][5].IsOccupied = true
		p, err = piece.NewPiece("wn")
		if err != nil {
			return err
		}
		gameBoard[7][6].Piece = p
		gameBoard[7][6].IsOccupied = true
		p, err = piece.NewPiece("wr")
		if err != nil {
			return err
		}
		gameBoard[7][7].Piece = p
		gameBoard[7][7].IsOccupied = true
		p, err = piece.NewPiece("wp")
		if err != nil {
			return err
		}
		for i := 0; i < 8; i++ {
			gameBoard[6][i].Piece = p
			gameBoard[6][i].IsOccupied = true
		}
	}
	return nil
}
