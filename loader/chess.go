package main

import (
	"github.com/notnil/chess"
	"strings"
	"fmt"
)

func ObjectifyGame(pgn string) *chess.Game {
    reader := strings.NewReader(pgn)
	match, err := chess.PGN(reader)
	if err != nil {
		fmt.Println("ERROR: Could not load PGN with notnil/chess")
	}
	game := chess.NewGame(match)
	//fmt.Print(game)
	return game
}