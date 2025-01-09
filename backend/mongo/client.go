package mongo

import (
	"context"
	"fmt"
	"strings"
	"errors"
	"time"
	"embed"

	"go.mongodb.org/mongo-driver/bson"
	_"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
)

//go:embed creds
var creds embed.FS

var uri string = readCreds();
var client *mongo.Client

func readCreds() string{
	data, err := creds.ReadFile("creds")
	if err != nil {
		panic("[ERROR] no creds file")
	}
	return strings.Split(string(data), "\n")[0]
}

func Init() {
	client = oneShotClient()
}

func MongoFindRandPuzzle() (*data.Puzzle, error){
	
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

	coll := client.Database("official").Collection("puzzles")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result []data.Puzzle
	filter := mongo.Pipeline{{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}} }

	cursor, err := coll.Aggregate(ctx,filter)
	if err != nil {
		fmt.Println("[ERROR] could not load random puzzle from database because:" , err)
		return nil, errors.New("could not load random puzzle from database")
	}
	if err = cursor.All(ctx, &result); err != nil {
		fmt.Println("[ERROR] could not extract random puzzle from result because:", err)
		return nil, errors.New("could not extract random puzzle from result")
	}

	if len(result) == 0 {return nil, errors.New("could not find any result")}
	
	fmt.Println("[SUCCESS] found random puzzle with id:", result[0].ID)

	return &result[0], nil
}

func oneShotClient() *mongo.Client {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("[SUCCESS] You successfully connected to MongoDB!")
	return client
}

func Destroy() {
	if err := client.Disconnect(context.TODO()); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("[SUCCESS] destroy Mongo client")
}

func IncrementSolved(pid string){

	coll := client.Database("official").Collection("puzzles")

	pipe := bson.D{
		{"$inc", bson.D{
			{"solvecount", 1},
		}},
	}
	filter := bson.D{{"id", pid}}

	_, err := coll.UpdateOne(context.TODO(), filter, pipe)
	if err != nil {
		fmt.Println("[ERROR] could not increment solvecount for",pid," because:" , err)
	}
}
func MarkAsSeen(pid string, uuid string) {
	
}