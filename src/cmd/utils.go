package cmd

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func GetMongoDatabase(databaseName string) *mongo.Database {
	databaseURL := getEnvironmentVariable("DATABASE_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(databaseName)
}
