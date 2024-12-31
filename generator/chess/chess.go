package chess

import (
	"github.com/notnil/chess"
	"strings"
	"fmt"
	"math/rand/v2"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
)

func ObjectifyGame(pgn string) *chess.Game {
    reader := strings.NewReader(pgn)
	match, err := chess.PGN(reader)
	if err != nil {
		fmt.Println("[ERROR] Could not load PGN with notnil/chess")
	}
	game := chess.NewGame(match)
	return game
}

func IsCheckmate(game *chess.Game) bool{
	return game.Method() == chess.Checkmate
}

func NavToMove(mvNum int, game *chess.Game) string{
	f := chess.NewGame()

	for mv:=0; mv<=(len(game.Moves()) - mvNum)-1; mv++{
		f.Move(game.Moves()[mv])
	}
	return f.FEN()
}

func GetFinalBestMoves(game *chess.Game) [config.BestMovesNumber]string{

	var bestmvs [config.BestMovesNumber]string
	j:=config.BestMovesNumber-1
	for i := len(game.Moves())-1; i>((len(game.Moves())-1) - config.BestMovesNumber); i-- {
		bestmvs[j] = game.Moves()[i].String()
		j--
	}
	return bestmvs
}
func JumpToRandPosition(game *chess.Game) (*chess.Game, int){
	
	x := rand.IntN(len(game.Moves()))
	fmt.Println("[SUCCESS] random position set: ", x)
	game.Move(game.Moves()[x])

	newG := chess.NewGame()
	for i:=0; i <= x-1; i++ {
		newG.Move(game.Moves()[i])
	}

	return newG, x
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


func IsGameEligible(game *chess.Game, randPos int) bool {

	//1) check moves number
	totalMoves := len(game.MoveHistory())
	if totalMoves < config.TolaratedNumberOfMoves {return false}
	
	//2) check if random move number picked is not out of bounds when running stockfish later
	if (totalMoves - randPos) < config.BestMovesNumber {return false}

	//3) match has ended either by black/white winning only
	if game.Method() != chess.Checkmate {return false}
	

	return true
}

func IsGameOver(game *chess.Game) bool{
	return game.Outcome() == chess.NoOutcome
}

func JumpToBeforeCheckmate(game *chess.Game) (*chess.Game, int){
	
	totalMoves := len(game.MoveHistory())
	x := totalMoves - config.BestMovesNumber
	newG := chess.NewGame()
	for i:=0; i <= x-1; i++ {
		newG.Move(game.Moves()[i])
	}

	return newG, x
}

func DetermineTurn(game *chess.Game) int{
	whosPlaying := 0;
	if game.Position().Turn().Name() == "Black" {
		whosPlaying = 0
	}else{
		whosPlaying = 1
	}
	return whosPlaying
}