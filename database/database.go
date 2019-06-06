package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ctx context.Context
var Client *mongo.Client
var err error

func Initialize() error {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	//TODO: Alterar para base compartilhada do squad
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://squad1:$quad1floripa@codenationsquad1-agrei.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return err
}
