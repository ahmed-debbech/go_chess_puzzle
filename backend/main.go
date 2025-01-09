package main

import (
	"fmt"
	"net/http"
	"embed"
	"io/fs"
	_"time"
	"github.com/ahmed-debbech/go_chess_puzzle/backend/logic"
	"github.com/ahmed-debbech/go_chess_puzzle/backend/ramstore"
)

//go:embed views/*.html
var html embed.FS

//go:embed views/assets/**
var assets embed.FS

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) *Page {

	body, err := html.ReadFile("views/" + title)
    if err != nil {
		fmt.Println("[ERROR] could not load page", title, err)
		return nil
	}

	return &Page{Title: title, Body: body}
}


func loadPuzzleHandle(w http.ResponseWriter, r *http.Request){
	if(r.Method != "GET"){http.Error(w, "Invalid", http.StatusMethodNotAllowed); return;}

	puzzle, err := logic.GetRandomPuzzle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	serial, err := logic.PuzzleToJson(*puzzle)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		
		w.Write([]byte(`{"error": "`+err.Error()+`"}`))
	}
	ramstore.SetToRamStore(puzzle)
	w.Header().Set("Content-Type", "application/json")
	w.Write(serial)
}

func solvedHandler(w http.ResponseWriter, r *http.Request){
	if(r.Method != "GET"){http.Error(w, "Invalid", http.StatusMethodNotAllowed); return;}
    
	query := r.URL.Query()
    puzzleId := query.Get("pid")
	hash := query.Get("h")

	solved := ramstore.CheckRamStore(puzzleId, hash)

	if solved {
		go logic.IncrementSolvedCounter(puzzleId)
		w.Write([]byte("true"))
	}else{
		w.Write([]byte("false"))
	}
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	p := loadPage("index.html")
	if p == nil { return }

	fmt.Fprintln(w,string(p.Body))
}


func main(){
	fmt.Println("Hello world")

	logic.ConnectDb()
	defer logic.StopDb()

    staticFiles, _ := fs.Sub(assets, "views/assets")

	http.Handle(
		"/assets/", 
		http.StripPrefix(
			"/assets/", 
			http.FileServer(http.FS(staticFiles)),
		),
	)
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/load", loadPuzzleHandle)
	http.HandleFunc("/solved", solvedHandler)


    fmt.Println(http.ListenAndServe(":5530", nil))
}