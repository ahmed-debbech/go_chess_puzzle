package chess

import (
	"github.com/notnil/chess"
	"strings"
	"fmt"
	"math/rand/v2"
)

func ObjectifyGame(pgn string) *chess.Game {
    reader := strings.NewReader(pgn)
	match, err := chess.PGN(reader)
	if err != nil {
		fmt.Println("[ERROR] Could not load PGN with notnil/chess")
	}
	game := chess.NewGame(match)
	//fmt.Print(game)
	return game
}

func JumpToRandPosition(game *chess.Game) (*chess.Game){
	
	x := rand.IntN(len(game.Moves()))
	fmt.Println("[SUCCESS] random position set: ", x)
	game.Move(game.Moves()[x])

	newG := chess.NewGame()
	for i:=0; i <= x-1; i++ {
		newG.Move(game.Moves()[i])
	}

	return newG
}

func GenerateFen(game *chess.Game) string{
	fmt.Println("[SUCCESS] generating FEN")
	return game.FEN()
}