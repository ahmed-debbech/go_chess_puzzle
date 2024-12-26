package routes

import (
	"fmt"
	"github.com/ahmed-debbech/go_chess_puzzle/storer/mongo"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"

)

func SaveToMongo(puzzle data.Puzzle) (string, bool){

	succ := true
	id := mongo.InsertPuzzle(puzzle)
	if id == "" {
		succ = false
	}
	return id ,succ
}

func rawToPuzzle(dat map[string]interface{}) (data.Puzzle){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[ERROR] something went wrong in converting types:", err)
			panic("")
		}
	}()

	strSlice := make([]string, len(dat["BestMoves"].([]interface{})))
	var strArray [config.BestMovesNumber]string
    for i, v := range dat["BestMoves"].([]interface{}) {
        strSlice[i], _ = v.(string)
    }	

	for i:=0;i<config.BestMovesNumber; i++ {
		strArray[i] = strSlice[i]
	}
	fmt.Println(strArray)
	puzzle := data.Puzzle{
		ID: dat["ID"].(string),
		FEN: dat["FEN"].(string),
		BestMoves: strArray,
		GenTime: dat["GenTime"].(string),
		SolveCount: int(dat["SolveCount"].(float64)),
		MatchLink: dat["MatchLink"].(string),
		SeenCount: int(dat["SeenCount"].(float64)),
		FirstSeenTime: dat["FirstSeenTime"].(string),
		CurrentPlayer: int(dat["CurrentPlayer"].(float64)),
	}
	fmt.Println("[SUCCESS] converting from raw to data.Puzzle")
	return puzzle
}