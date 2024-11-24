package mongo

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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