package data;

import (
	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
)

type Puzzle struct {
	ID string
	FEN string
	BestMove [config.BestMovesNumber]string
	GenTime string
	SolveTime string
	MatchLink string
	SeenCount string
	FirstSeenTime string
}