package logic

import (
	"os"
	"strconv"
	"math/rand/v2"
	//"fmt"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/utils"
)

func LookupMatch(path string, max int) (string, int){
	x  := rand.IntN(max)
	x = 135311
	data, err := os.ReadFile(utils.EndItWithSlash(path)+strconv.Itoa(x))

	for err != nil {
		x = rand.IntN(max)
		x = 135311
		//fmt.Println(x)
		data, err = os.ReadFile(utils.EndItWithSlash(path)+strconv.Itoa(x))
	}

	return string(data), x
}
