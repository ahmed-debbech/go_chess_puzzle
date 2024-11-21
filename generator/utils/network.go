package utils

import(
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
	"time"
	"fmt"
	"net/http"
	"bytes"

)

func SendToStore(puzzle data.Puzzle) (error){

	jsonPuzzle, err := puzzle.ToJson()
	if err != nil {
		fmt.Println("[ERROR] could not serialize to json data")
		return err
	}else{
		fmt.Println("[SUCCESS] puzzle generated to json ", string(jsonPuzzle))
		req, err := http.NewRequest("POST", "http://localhost:3000/api/accept", bytes.NewReader(jsonPuzzle))
		client := http.Client{Timeout: 10 * time.Second}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return err
		}else{
			if res.Status == "200 OK" {
				fmt.Println("[SUCCESS] storer responded with 200 OK - Puzzle uploaded to store")
				return nil
			}else{
				fmt.Println("[ERROR] something when wrong when uploading to storer") 
				return err
			}
		}
	}
}