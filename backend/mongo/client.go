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

var uri , dbname  = readCreds();
var client *mongo.Client

func readCreds() (string, string){
	data, err := creds.ReadFile("creds")
	if err != nil {
		panic("[ERROR] no creds file")
	}
	return strings.Split(string(data), "\n")[0], 
	strings.Split(string(data), "\n")[1]
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

	coll := client.Database(dbname).Collection("puzzles")

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

	coll := client.Database(dbname).Collection("puzzles")

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
	coll := client.Database(dbname).Collection("puzzles")

	pipe := bson.D{{"$addToSet", bson.D{{"seencount", uuid}},}}

	filter := bson.D{{"id", pid}}

	_, err := coll.UpdateOne(context.TODO(), filter, pipe)
	if err != nil {
		fmt.Println("[ERROR] could not increment solvecount for",pid," because:" , err)
	}
}

func GetUniquePlayers() (int32, error){
	
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

	coll := client.Database(dbname).Collection("puzzles")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result []bson.M
	filter := mongo.Pipeline{
		{{"$unwind", "$seencount"}},
		{{"$group", bson.D{{"_id", "$seencount"}}}},
		{{"$count", "uniqueUUIDs"}},
	}

	cursor, err := coll.Aggregate(ctx,filter)
	if err != nil {
		fmt.Println("[ERROR] could not get value from GetUniquePlayers()" , err)
		return -1, errors.New("could not get value from GetUniquePlayers()")
	}
	if err = cursor.All(ctx, &result); err != nil {
		fmt.Println("[ERROR] could not extract random puzzle from result because:", err)
		return -1, errors.New("could not extract random puzzle from result")
	}

	if len(result) == 0 {return -1, errors.New("could not find any result")}
	
	return result[0]["uniqueUUIDs"].(int32), nil
}

func GetTotalNumberOfSolves() (int32, error){
	
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

	coll := client.Database(dbname).Collection("puzzles")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result []bson.M
	filter := []bson.M{
		{
			"$group": bson.M{
				"_id":        "",
				"solvecount": bson.M{"$sum": "$solvecount"},
			},
		},
	}
	cursor, err := coll.Aggregate(ctx,filter)
	if err != nil {
		fmt.Println("[ERROR] could not get value from GetTotalNumberOfSolves()" , err)
		return -1, errors.New("could not get value from GetTotalNumberOfSolves()")
	}
	if err = cursor.All(ctx, &result); err != nil {
		fmt.Println("[ERROR] could not extract random puzzle from result because:", err)
		return -1, errors.New("could not extract random puzzle from result")
	}

	if len(result) == 0 {return -1, errors.New("could not find any result")}
	
	return result[0]["solvecount"].(int32), nil
}