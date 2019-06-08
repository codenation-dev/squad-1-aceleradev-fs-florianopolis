package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Context context.Context
var Client *mongo.Client
var err error
var dbName = "projeto-final"

func Initialize() error {
	Context = context.TODO()
	url := os.Getenv("URL_MONGO")
	Client, err = mongo.Connect(Context, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return err
}

func GetCollection(name string) *mongo.Collection {
	collection := Client.Database(dbName).Collection(name)
	return collection
}
