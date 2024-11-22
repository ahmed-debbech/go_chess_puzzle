package main

import (
    "fmt"
    "os"
    "net/http"
    "github.com/ahmed-debbech/go_chess_puzzle/storer/routes"
    "github.com/ahmed-debbech/go_chess_puzzle/storer/mongo"
)

func main() {
    fmt.Println("Hello, Storer!")

    mongo.Client()

    http.Handle("/", http.HandlerFunc(routes.Root))
    http.Handle("/accept", http.HandlerFunc(routes.Accept))

    err := http.ListenAndServe(":5500", nil)
    if err != nil {
        fmt.Println("[ERROR] could not start web server")
        os.Exit(1) 
    }
    fmt.Println("[SUCCESS] Server Up and Running")
}