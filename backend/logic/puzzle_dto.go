package logic

import (
	"strconv"
	"encoding/json"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
)

type PuzzleDto struct {
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

func (p PuzzleDto) String() string{
	return p.ID + " FEN: " + p.FEN + " BestMove: " + strconv.Itoa(len(p.BestMoves)) + " GenTime: " + p.GenTime + " CurrentPlayer: "+ strconv.Itoa(p.CurrentPlayer)  +" SolveCount: "  + strconv.Itoa(p.SolveCount) + " MatchLink: " + p.MatchLink + " SeenCount: " + strconv.Itoa(p.SeenCount) + " FirstSeenTime: " + p.FirstSeenTime;
}

func fromPuzzleDao(puzzleDao *data.Puzzle) *PuzzleDto{
	return &PuzzleDto{
		ID: puzzleDao.ID,
		FEN: puzzleDao.FEN,
		BestMoves: puzzleDao.BestMoves,
		GenTime: puzzleDao.GenTime,
		CurrentPlayer: puzzleDao.CurrentPlayer,
		SolveCount: puzzleDao.SolveCount,
		MatchLink: puzzleDao.MatchLink,
		SeenCount: len(puzzleDao.SeenCount),
		FirstSeenTime: puzzleDao.FirstSeenTime,
	}
}

func (p PuzzleDto) ToJson() ([]byte, error){
	return json.Marshal(p)
}