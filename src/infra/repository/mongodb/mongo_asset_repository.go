package repository

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "assets"

type MongoAssetRepository struct {
	collection *mongo.Collection
}

func (m MongoAssetRepository) Insert(asset entities.Asset) (entities.Asset, error) {
	_, err := m.collection.InsertOne(nil, asset)
	return asset, err
}

func (m MongoAssetRepository) GetLastPosition() (int, error) {
	lastAssetOfList := entities.Asset{}
	opts := options.FindOne().SetSort(bson.D{{Key: "order", Value: -1}})
	m.collection.FindOne(nil, bson.D{}, opts).Decode(&lastAssetOfList)

	return lastAssetOfList.Order, nil
}

func (m MongoAssetRepository) CheckIfAssetExists(code string) (bool, error) {
	asset := entities.Asset{}
	m.collection.FindOne(nil, bson.D{{Key: "code", Value: code}}).Decode(&asset)

	return asset.Code != "", nil
}

func NewAssetRepository(database mongo.Database) MongoAssetRepository {
	return MongoAssetRepository{
		collection: database.Collection(COLLECTION_NAME),
	}
}
