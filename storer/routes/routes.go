package routes


import (
	"fmt"
	"net/http"
	"encoding/json"
	"bufio"
	"io"
	"errors"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
)


func Root(w http.ResponseWriter, r *http.Request){
	fmt.Println("[HIT] route /")
}

func Accept(w http.ResponseWriter, r *http.Request){

	if allowedMethod(r, http.MethodPost) != nil {http.Error(w, "", http.StatusMethodNotAllowed); return}

	fmt.Println("[HIT] route /accept")

	w.Header().Set("Content-Type", "application/json")

	body, err := getBody(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return		
	}

	var dat map[string]interface{}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		fmt.Println("[ERROR] could not serialize object coming from client")
	  	http.Error(w, "Could not serialize", http.StatusInternalServerError)
	  	return
	}
	puzzle := data.Puzzle{
		ID: dat["ID"].(string),
		FEN: dat["FEN"].(string),
	}

	rb, _ := json.Marshal(puzzle)
	w.Write(rb)
}


func getBody(r io.Reader) ([]byte, error) {
	
	bb := make([]byte, 0)

	reader := bufio.NewReader(r)
	for {
		char, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("[ERROR] could not read request body")
			return nil, err
		}
		bb = append(bb, char)
	}
	return bb, nil
}

func allowedMethod(r *http.Request, method string) error{
	if r.Method != method { 
		fmt.Println("[ERROR] StatusMethodNotAllowed")
		return errors.New("StatusMethodNotAllowed")
	}
	return nil
}