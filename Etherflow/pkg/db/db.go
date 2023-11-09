package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewMongoClient() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func Init() {

	client, err := NewMongoClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"address": 1},
		Options: options.Index().SetUnique(true),
	}

	collection := client.Database("Etherflow").Collection("Users")
	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func Insert(data interface{}) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := NewMongoClient()

	defer client.Disconnect(ctx)

	collection := client.Database("Etherflow").Collection("Users")

	collection.InsertOne(context.Background(), data)

}
