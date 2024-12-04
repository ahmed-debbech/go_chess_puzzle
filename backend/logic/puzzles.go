package logic

import (
	"errors"
	"github.com/ahmed-debbech/go_chess_puzzle/backend/mongo"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
)


func GetRandomPuzzle() (*data.Puzzle, error){
	dat, err := mongo.MongoFindRandPuzzle()
	if err != nil {
		return &data.Puzzle{}, errors.New("Could not find a random puzzle.")
	}
	return dat, nil
}