package main

import (
	"fmt"
	"os"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/chess"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/logic"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/utils"
)


func main() {
	
	if len(os.Args) <= 2{
		fmt.Println("NO ARGS!!")
		os.Exit(1)
	} 

	max, yes := utils.IsNumber(os.Args[1]);
	if yes == false {
		fmt.Println("ERROR: Set the maximum number correctly")
		os.Exit(1)
	}

	if !utils.IsDirectoryExist(os.Args[2]) {
		fmt.Println("ERROR: Not a valid directory to read from.")
		os.Exit(1)
	}

	fmt.Println("Looking for a game match from this directory: ", os.Args[2])

	match_content, id := logic.LookupMatch(os.Args[2], max)
	//fmt.Println(match_content)
	fmt.Println("[SUCCESS] MATCH FOUND! ID: ", id)

	game := chess.ObjectifyGame(match_content)
	fmt.Println(game.Position().Board().Draw())


}