package routes


import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request){
	fmt.Println("[HIT] route /")
}

func Accept(w http.ResponseWriter, r *http.Request){
	fmt.Println("[HIT] route /accept")
	
}