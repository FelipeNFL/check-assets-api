package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

const COLLECTION_NAME = "assets"
const DEFAULT_TIMEOUT = 10

type MongoAssetRepository struct {
	collection *mongo.Collection
}

func (m MongoAssetRepository) getContext(timeout_optional ...int) (context.Context, context.CancelFunc) {
	timeout := DEFAULT_TIMEOUT

	if len(timeout_optional) > 0 {
		timeout = timeout_optional[0]
	}

	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}

func (m MongoAssetRepository) Insert(asset entities.Asset) (entities.Asset, error) {
	ctx, cancel := m.getContext()
	defer cancel()

	_, err := m.collection.InsertOne(ctx, asset)

	return asset, err
}

func (m MongoAssetRepository) GetLastPosition() (int, error) {
	ctx, cancel := m.getContext()
	defer cancel()

	lastAssetOfList := entities.Asset{}
	opts := options.FindOne().SetSort(bson.D{{Key: "order", Value: -1}})
	m.collection.FindOne(ctx, bson.D{}, opts).Decode(&lastAssetOfList)

	return lastAssetOfList.Order, nil
}

func (m MongoAssetRepository) CheckIfAssetExists(code string) (bool, error) {
	ctx, cancel := m.getContext()
	defer cancel()

	asset := entities.Asset{}
	m.collection.FindOne(ctx, bson.D{{Key: "code", Value: code}}).Decode(&asset)

	return asset.Code != "", nil
}

func (m MongoAssetRepository) GetAll() ([]entities.Asset, error) {
	ctx, cancel := m.getContext()
	defer cancel()

	cursor, err := m.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	assets := []entities.Asset{}

	if err = cursor.All(ctx, &assets); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return assets, nil
}

func NewAssetRepository(database mongo.Database) MongoAssetRepository {
	return MongoAssetRepository{
		collection: database.Collection(COLLECTION_NAME),
	}
}
