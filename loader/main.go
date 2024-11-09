package main

import (
	"fmt"
	"os"
	"math/rand/v2"
	"strconv"
)

func IsMatchFound(path string, max int) (string, int){
	x  := rand.IntN(max)
	data, err := os.ReadFile(path+strconv.Itoa(x))

	for err != nil {
		x = rand.IntN(max)
		data, err = os.ReadFile(path+strconv.Itoa(x))
	}

	return string(data), x
}


func main() {
	
	if len(os.Args) <= 1{
		fmt.Println("NO ARGS!!")
		os.Exit(1)
	} 
	max, err := strconv.Atoi(os.Args[1]);
	if err != nil {
		fmt.Println("ERROR: Set the maximum number correctly")
		os.Exit(1)
	}
	
	fmt.Println("Looking for a game match from this directory: ", os.Args[2])

	match_content, id := IsMatchFound(os.Args[2], max)
	//fmt.Println(match_content)
	fmt.Println("MATCH FOUND! ID: ", id)

	game := ObjectifyGame(match_content)
	fmt.Println(game.Position().Board().Draw())


}