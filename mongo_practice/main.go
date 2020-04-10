package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	readCollection()
}

func initMongo() {

	var err error

	// var cancel func()

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel() // defer는 함수가 종료될 때 에러가 아니라면 가장 마지막으로 실행됨.

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

func insertData() {

	user := Model.User{
		Id:    primitive.NewObjectID(),
		Name:  "asdasd",
		Email: "asdas@asdas",
	}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}

	resId := result.InsertedID.(primitive.ObjectID)

	fmt.Println("생성된 데이터 id : ", resId.Hex())
}

func readCollection() {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		println(err)
	}

	for cursor.Next(ctx) {
		data := Model.User{}
		cursor.Decode(&data)
		println(data.Name)
	}
}
