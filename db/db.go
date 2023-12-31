package db

import (
	"context"
	"log"
	"words-bot/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClientOptions() *options.ClientOptions {
	envs, err := utils.GetEnvs()
	if err != nil {
		log.Fatal(err)
	}

	dburi := envs.DbUri

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(dburi).
		SetServerAPIOptions(serverAPIOptions)

	return clientOptions
}

func GetCollection(collection string) (*mongo.Collection, error) {
	client := GetMongoClient()

	return client.Database("words").Collection(collection), nil
}

func InitDb() {
	clientOptions := GetClientOptions()

	newClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		client = newClient
	}
}

func GetMongoClient() mongo.Client {
	return *client
}
