package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
)

const (
	dbName = "testMongo"
)

func main() {
	initMongo()
	insertData()

}

func initMongo() {

	var err error
	var cancel func()
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
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

func insertData() {
	res, err := collection.InsertOne(ctx, bson.M{"name": "jhj", "id": "kk", "asda": "jnkjk"})

	if err != nil {
		panic(err)
	}

	id := res.InsertedID

	println("들어간 데이타 : ", id)
}
