package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	Model "Projects/first_golang/mongo_practice/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoUrl   string
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
	// readCollection()

}

func initMongo() {

	var err error
	var cancel func()

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getMongoUrl()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUrl,
	))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection("User")

	println("몽고디비에 연결되었습니다.")
}

func getMongoUrl() {
	txtFile, _ := os.Open("configs.txt")

	scanner := bufio.NewScanner(txtFile)

	for scanner.Scan() {
		line := scanner.Text()
		mongoUrl = line
		break
	}
}

func readCollection() {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		println(err)
	}

	// resultList := []Model.User{}

	for cursor.Next(ctx) {
		data := Model.User{}
		cursor.Decode(&data)
		println(data.Name)
	}
}

func insertData() {

	user := Model.User{
		Id:    primitive.NewObjectID(),
		Name:  "asdasd",
		Email: "asdas@asdas",
	}

	result, _ := collection.InsertOne(ctx, user)
	// if err != nil {
	// 	println(err)
	// 	panic(err)
	// }

	resId := result.InsertedID.(primitive.ObjectID).String()
	println("생성된 데이터 : ", resId)

}
