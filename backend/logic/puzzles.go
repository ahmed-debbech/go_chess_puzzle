package logic

import (
	"errors"
	"github.com/ahmed-debbech/go_chess_puzzle/backend/mongo"
)


func GetRandomPuzzle() (*PuzzleDto, error){
	dat, err := mongo.MongoFindRandPuzzle()
	if err != nil {
		return &PuzzleDto{}, errors.New("Could not find a random puzzle.")
	}
	pdto := fromPuzzleDao(dat)
	return pdto, nil
}

func PuzzleToJson(puzzle PuzzleDto) ([]byte, error){
	dat, err := puzzle.ToJson()
	if err != nil {
		return []byte{}, errors.New("Could serialize puzzle to JSON")
	}
	return dat, nil
}

func IncrementSolvedCounter(puzzleId string) {
	mongo.IncrementSolved(puzzleId)
}

func MarkPuzzleAsSeen(pid string, uuid string){
	mongo.MarkAsSeen(pid, uuid)
}