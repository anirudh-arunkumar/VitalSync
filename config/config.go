package config

import(
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func connectDB() {
	client := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error

	Client, err = mongo.Connect(context.TODO(), client)

	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Mongo")
}