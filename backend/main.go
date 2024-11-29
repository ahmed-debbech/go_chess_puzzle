package main

import (
	"os"
	"fmt"
	"net/http"
	"html/template"
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


func handler(w http.ResponseWriter, r *http.Request) {

	p := loadPage("index.html")
	if p == nil { return }

	t, err := template.ParseFiles(getPath("index.html"))
	if err != nil { fmt.Println("[ERROR] could not render page", err); return }
	t.Execute(w, p)
}

func main(){
	fmt.Println("Hello world")

	http.HandleFunc("/", handler)
	http.Handle(
		"/assets/", 
		http.StripPrefix(
			"/assets/", 
			http.FileServer(http.Dir("./views/assets/")),
		),
	)

    fmt.Println(http.ListenAndServe(":5530", nil))
}