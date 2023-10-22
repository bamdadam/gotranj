package chess

import (
	"github.com/bamdadam/gotranj/internal/chess/piece"
	"github.com/bamdadam/gotranj/internal/chess/tile"
)

func FillTilesWithPieces(isBlack bool, gameBoard [][8]*tile.Tile) ([]piece.Piece, error) {
	pcs := []piece.Piece{}
	if isBlack {
		p, err := piece.NewPiece("br", 0, 0)
		if err != nil {
			return nil, err
		}
		gameBoard[0][0].Piece = p
		gameBoard[0][0].IsOccupied = true
		pcs = append(pcs, p)
		p, err = piece.NewPiece("bn", 0, 1)
		if err != nil {
			return nil, err
		}
		gameBoard[0][1].Piece = p
		gameBoard[0][1].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("bb", 0, 2)
		if err != nil {
			return nil, err
		}
		gameBoard[0][2].Piece = p
		gameBoard[0][2].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("bq", 0, 3)
		if err != nil {
			return nil, err
		}
		gameBoard[0][3].Piece = p
		gameBoard[0][3].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("bk", 0, 4)
		if err != nil {
			return nil, err
		}
		gameBoard[0][4].Piece = p
		gameBoard[0][4].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("bb", 0, 5)
		if err != nil {
			return nil, err
		}
		gameBoard[0][5].Piece = p
		gameBoard[0][5].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("bn", 0, 6)
		if err != nil {
			return nil, err
		}
		gameBoard[0][6].Piece = p
		gameBoard[0][6].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("br", 0, 7)
		if err != nil {
			return nil, err
		}
		gameBoard[0][7].Piece = p
		gameBoard[0][7].IsOccupied = true
		pcs = append(pcs, p)

		for i := 0; i < 8; i++ {
			p, err = piece.NewPiece("bp", 1, i)
			if err != nil {
				return nil, err
			}
			gameBoard[1][i].Piece = p
			gameBoard[1][i].IsOccupied = true
			pcs = append(pcs, p)

		}
	} else {
		p, err := piece.NewPiece("wr", 7, 0)
		if err != nil {
			return nil, err
		}
		gameBoard[7][0].Piece = p
		gameBoard[7][0].IsOccupied = true
		pcs = append(pcs, p)
		p, err = piece.NewPiece("wn", 7, 1)
		if err != nil {
			return nil, err
		}
		gameBoard[7][1].Piece = p
		gameBoard[7][1].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("wb", 7, 2)
		if err != nil {
			return nil, err
		}
		gameBoard[7][2].Piece = p
		gameBoard[7][2].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("wq", 7, 3)
		if err != nil {
			return nil, err
		}
		gameBoard[7][3].Piece = p
		gameBoard[7][3].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("wk", 7, 4)
		if err != nil {
			return nil, err
		}
		gameBoard[7][4].Piece = p
		gameBoard[7][4].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("wb", 7, 5)
		if err != nil {
			return nil, err
		}
		gameBoard[7][5].Piece = p
		gameBoard[7][5].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("wn", 7, 6)
		if err != nil {
			return nil, err
		}
		gameBoard[7][6].Piece = p
		gameBoard[7][6].IsOccupied = true
		pcs = append(pcs, p)

		p, err = piece.NewPiece("wr", 7, 7)
		if err != nil {
			return nil, err
		}
		gameBoard[7][7].Piece = p
		gameBoard[7][7].IsOccupied = true
		pcs = append(pcs, p)

		for i := 0; i < 8; i++ {
			p, err = piece.NewPiece("wp", 6, i)
			if err != nil {
				return nil, err
			}
			gameBoard[6][i].Piece = p
			gameBoard[6][i].IsOccupied = true
			pcs = append(pcs, p)

		}
	}
	return pcs, nil
}
