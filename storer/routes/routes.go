package routes


import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)

type Puzzle struct{

}

func Root(w http.ResponseWriter, r *http.Request){
	fmt.Println("[HIT] route /")
}

func Accept(w http.ResponseWriter, r *http.Request){
	fmt.Println("[HIT] route /accept")
	io.WriteString(w, "Eee\n")
	puzzle := Puzzle{}
	err := json.NewDecoder(r.Body).Decode(&puzzle)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}