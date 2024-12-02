package mongo

import (
	"context"
	"fmt"
	"os"
	"strings"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	_"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
)

var uri string = readCreds();
var client *mongo.Client

func readCreds() string{
	data, err := os.ReadFile("mongo/creds")
	if err != nil {
		panic("[ERROR] no creds file")
	}
	return strings.Split(string(data), "\n")[0]
}

func Init() {
	client = oneShotClient()
}

func MongoFindRandPuzzle() (data.Puzzle, error){
	coll := client.Database("official").Collection("puzzles")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var result data.Puzzle
	filter := mongo.Pipeline{{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}} }

	cursor, err := coll.Aggregate(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("[ERROR] the query succeeded but it seems like it is empty")
			return result, errors.New("No Puzzles found")
		}
		fmt.Println("[ERROR] something went wrong while finding a random puzzle:", err)
		return result, errors.New("Could not find random puzzle")
	}
	defer cursor.Close(ctx)

	/*for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}*/

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	fmt.Printf("Average price of %v \n", result["_id"])

	return result, nil
}

func oneShotClient() *mongo.Client {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
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