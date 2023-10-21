package main

import (
	"fmt"

	"github.com/bamdadam/gotranj/internal/chess/game"
	"github.com/bamdadam/gotranj/internal/player"
)

func main() {
	player1 := player.Player{}
	player2 := player.Player{}
	game := game.NewGame(player1, player2)
	fmt.Println(game.Board)
}
