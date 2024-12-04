package main

import (
	"os"
	"fmt"
	"net/http"
	"html/template"
	"github.com/ahmed-debbech/go_chess_puzzle/backend/logic"
)

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) *Page {
    pwd, err := os.Getwd()

	body, err := os.ReadFile(pwd + "/views/" + title)
    if err != nil {
		fmt.Println("[ERROR] could not load page", title, err)
		return nil
	}

	return &Page{Title: title, Body: body}
}

func getPath(title string) string {
	pwd, _ := os.Getwd()
	return pwd + "/views/" + title
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(serial)
}
func rootHandle(w http.ResponseWriter, r *http.Request) {
	p := loadPage("index.html")
	if p == nil { return }

	t, err := template.ParseFiles(getPath("index.html"))
	if err != nil { fmt.Println("[ERROR] could not render page", err); return }
	t.Execute(w, p)
}

func main(){
	fmt.Println("Hello world")


	logic.ConnectDb()
	defer logic.StopDb()

	http.Handle(
		"/assets/", 
		http.StripPrefix(
			"/assets/", 
			http.FileServer(http.Dir("./views/assets/")),
		),
	)
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/load", loadPuzzleHandle)


    fmt.Println(http.ListenAndServe(":5530", nil))
}