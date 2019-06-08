package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Context context.Context
var Client *mongo.Client
var err error
var dbName = "projeto-final"

func Initialize() error {
	Context = context.TODO()
	//TODO: Alterar para obter URL a partir de variável de ambiente
	Client, err = mongo.Connect(Context, options.Client().ApplyURI("mongodb+srv://squad1:$quad1floripa@codenationsquad1-agrei.mongodb.net/test?retryWrites=true&w=majority"))
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
