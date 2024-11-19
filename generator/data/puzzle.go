package data;

import (
	"strconv"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
)

type Puzzle struct {
	ID string
	FEN string
	BestMoves [config.BestMovesNumber]string
	GenTime string
	SolveTime string
	MatchLink string
	SeenCount int
	FirstSeenTime string
}

func (p Puzzle) String() string{
	return p.ID + " FEN: " + p.FEN + " BestMove: " + strconv.Itoa(len(p.BestMoves)) + " GenTime: " + p.GenTime + " SolveTime: " + p.SolveTime + " MatchLink: " + p.MatchLink + " SeenCount: " +strconv.Itoa(p.SeenCount) + " FirstSeenTime: " + p.FirstSeenTime;
}