package routes


import (
	"fmt"
	"net/http"
	"encoding/json"

)


type Resp struct {
	Body interface{}
}

func Root(w http.ResponseWriter, r *http.Request){
	fmt.Println("[HIT] route /")
}

func Accept(w http.ResponseWriter, r *http.Request){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[ERROR] something went wrong!", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

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

	puzzle := rawToPuzzle(dat)
	
	id, succ := SaveToMongo(puzzle)

	res := Resp{
		Body: map[string]interface{}{
			"success" : succ,
			"_id" : id,
		},
	}
	rb, _ := json.Marshal(res)
	w.Write(rb)
}