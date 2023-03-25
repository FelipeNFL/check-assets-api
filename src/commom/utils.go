package commom

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func GetMongoDatabase(databaseName string) *mongo.Database {
	databaseURL := GetEnvironmentVariable("DATABASE_URL")

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
