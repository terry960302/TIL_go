package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

const (
	dbName = "testMongo"
)

func main() {
	initMongo()

}

func initMongo() {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://terry960302:ljk041180@ritier-joagf.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

}

func readCollection() {
	collection = client.Database(dbName).Collection("User")

}
