package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInstance() (client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoDB := "mongodb://localhost:27017"
	opts := options.Client().ApplyURI(mongoDB)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("connected database")
	return 
}

var Client *mongo.Client = DatabaseInstance()

func openCollection(client *mongo.Client, collectionName string) (collection *mongo.Collection) {
	collection = client.Database("gorestaurant").Collection(collectionName)
	return 
}
