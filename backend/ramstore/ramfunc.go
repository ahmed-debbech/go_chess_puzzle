package ramstore

import (
	"fmt"
	"github.com/ahmed-debbech/go_chess_puzzle/backend/logic"
)

func SetToRamStore(puzzle *logic.PuzzleDto){
	fmt.Println("Setting puzzle to RamStore")
	GetRamStoreInstance()
	Set(puzzle.ID, Calculate(puzzle.ID, puzzle.BestMoves))
}

func CheckRamStore(pid string, hash string) bool{
	fmt.Println("Checking if puzzle",pid, "is solved correctly")

	GetRamStoreInstance()
	hashFromRS, err := Get(pid)
	if err != nil {
		return false
	}
	fmt.Println("puzzle ",pid,"found in Ramstore")
	if hash == hashFromRS {
		fmt.Println("puzzle",pid,"has been solved correctly")
		Delete(pid)
		return true;
	}
	return false
}