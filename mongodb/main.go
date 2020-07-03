package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var host string = "mongodb://root:changeme@localhost:27017"
var collection *mongo.Collection

func insertOne() {
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "World"})
	if err != nil {
		log.Println(err)
	}
	id := res.InsertedID
	log.Println("id:", id)
}

func findOne() {
	var result interface{}
	err := collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

func findAll() {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			hello string
		}{}

		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}
		log.Println("result:", result)

		// use raw bson
		raw := cur.Current

		log.Println("raw: ", raw)
	}
	if err != nil {
		log.Println("EOF error:", err)
	}
}

func main() {
	log.Println("Start server")
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		log.Panicln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Panicln(err)
	}

	collection = client.Database("dv4test").Collection("blogs")

	insertOne()
	insertOne()
	insertOne()

	findOne()

	findAll()
}
