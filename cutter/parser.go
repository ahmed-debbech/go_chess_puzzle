package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math/rand/v2"
)

type OnePgnGame struct {
	Tags []byte
	Moves []byte
}

func prepareDataToSave(s OnePgnGame) ([]byte, int){
	id := rand.IntN(4000000000)
	var d []byte = make([]byte, 0)
	d = append(d,s.Tags...) 
	d = append(d, []byte{0x0A}...)
	d = append(d, []byte{0x0A}...)
	d = append(d,s.Moves...) 
	d = append(d, []byte{0x0A}...)
	d = append(d, []byte{0x0A}...)
	var dump []byte = make([]byte, int(len(s.Tags) + len(s.Moves)) + 2)
	dump = append(d)
	return dump, id
}

func main() {

	if len(os.Args) <= 1{
		fmt.Println("NO ARGS!!")
		os.Exit(1)
	} 

	file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	var f int8 = 1
	opg := OnePgnGame{
		Tags : make([]byte, 0),
		Moves : make([]byte, 0),
	}
	var counter int = 0;
    for scanner.Scan() {
		if scanner.Text() == "" {
			f++
			if f == 3 {

				//TODO add a way to index files incrementatly rather then randomly
				b, id := prepareDataToSave(opg)
				err = os.WriteFile("games/"+strconv.Itoa(id), b, 0644)
				if err != nil{
					fmt.Println(err)
					panic("could not write new file to store a game's pgn")	
				}

				f = 1
				opg = OnePgnGame{
					Tags : make([]byte, 0),
					Moves : make([]byte, 0),
				}
			}
			continue
		}
		if f == 1 {
			opg.Tags = append(opg.Tags, []byte(scanner.Text())...)
		}
		if f == 2 {
			opg.Moves = append(opg.Moves, []byte(scanner.Text())...)
		}
		fmt.Println("counter: ", counter)
		counter++
	}

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
	//fmt.Printf("Size %v\n", len(pgns))

}