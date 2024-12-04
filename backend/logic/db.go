package logic


import (
	"github.com/ahmed-debbech/go_chess_puzzle/backend/mongo"
)

func ConnectDb(){
	mongo.Init()
}

func StopDb(){
    defer mongo.Destroy()
}