package fileman

import (

	"strconv"
	"fmt"
	"os"
	"bufio"
)

type OnePgnGame struct {
	Tags []byte
	Moves []byte
}


func exportMatch(s OnePgnGame) ([]byte, string){
	var id string
	var c int 
	match_count, err := os.ReadFile("match_count")
	if err != nil {
		match_count = []byte("1")
	}else{
		c, _ = strconv.Atoi(string(match_count))
		c++
		id = strconv.Itoa(c)
	}
	err = os.WriteFile("match_count", []byte(strconv.Itoa(c)) , 0644)
	if err != nil {
		panic("could not write new match_count conter in a file")
	}

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

func endItWithSlash(path string) string{
	if path[len(path)-1] == '/' {
		return path
	}
	v := []byte(path)
	v = append(v, '/')
	return string(v)
}

func Scan(file *os.File, dir string){
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
				b, id := exportMatch(opg)
				err := os.WriteFile(endItWithSlash(dir) + id, b, 0644)
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
			opg.Tags = append(opg.Tags, []byte(scanner.Text()+string('\n'))...)
		}
		if f == 2 {
			opg.Moves = append(opg.Moves, []byte(scanner.Text())...)
		}
		counter++
	}

    if err := scanner.Err(); err != nil {
        fmt.Println("[ERROR] Error reading file: ", err)
		panic("[ERROR] Error reading file.")
    }
	fmt.Println("[SUCCESS] Done parsing all matches from the pgn file and exporting them")

}