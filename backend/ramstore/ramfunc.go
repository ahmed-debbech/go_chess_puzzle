package ramstore

import (
	"fmt"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
)

func SetToRamStore(puzzle *data.Puzzle){
	fmt.Println("Setting puzzle to RamStore")
	GetRamStoreInstance()
	Set(puzzle.ID, Calculate(puzzle.ID, puzzle.BestMoves))
	
}