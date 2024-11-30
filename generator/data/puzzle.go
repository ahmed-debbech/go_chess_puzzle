package data;

import (
	"strconv"
	"encoding/json"

	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
)

type Puzzle struct {
	ID string
	FEN string
	BestMoves [config.BestMovesNumber]string
	GenTime string
	CurrentPlayer int
	SolveCount int
	MatchLink string
	SeenCount int
	FirstSeenTime string
}

func (p Puzzle) String() string{
	return p.ID + " FEN: " + p.FEN + " BestMove: " + strconv.Itoa(len(p.BestMoves)) + " GenTime: " + p.GenTime + " CurrentPlayer: "+ strconv.Itoa(p.CurrentPlayer)  +" SolveCount: " + strconv.Itoa(p.SolveCount) + " MatchLink: " + p.MatchLink + " SeenCount: " +strconv.Itoa(p.SeenCount) + " FirstSeenTime: " + p.FirstSeenTime;
}

func (p Puzzle) ToJson() ([]byte, error){
	return json.Marshal(p)
}