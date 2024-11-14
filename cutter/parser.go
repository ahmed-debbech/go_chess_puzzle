package main


import (
	"fmt"
	"os"
	"github.com/ahmed-debbech/go_chess_puzzle/cutter/fileman"
)

func main() {

	if len(os.Args) <= 2 {
		fmt.Println("NOT ENOUGH ARGS!!")
		os.Exit(1)
	} 

	file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("[ERROR] Error opening file to parse")
        os.Exit(1)
    }
    defer file.Close()

	dir, err := os.Open(os.Args[2])
    if err != nil {
        fmt.Println("[ERROR] Exporting directory does not exist")
        os.Exit(1)
    }
    defer dir.Close()

	fileman.Scan(file, os.Args[2])
	//fmt.Printf("Size %v\n", len(pgns))

}