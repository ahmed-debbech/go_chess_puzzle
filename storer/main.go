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

    mongo.Init()
    defer mongo.Destroy()

    http.Handle("/", http.HandlerFunc(routes.Root))
    http.Handle("/accept", http.HandlerFunc(routes.Accept))

    fmt.Println("[SUCCESS] Server Up and Running")
    
    err := http.ListenAndServe(":5503", nil)
    if err != nil {
        fmt.Println("[ERROR] could not start web server")
        os.Exit(1) 
    }
}