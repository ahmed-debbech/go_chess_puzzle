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

func MakeMoveAndFEN(game *chess.Game, move string) string{

	p  := game.Position();
	uci := chess.UCINotation{}

	mv, err := uci.Decode(p, move);
	if err != nil {
		fmt.Println("[ERROR] move from SF not clear for our program")
		return ""
	}

	if err := game.Move(mv); err != nil {
		fmt.Println("[ERROR] could not make move: ", err)
	}
	return game.FEN()
}