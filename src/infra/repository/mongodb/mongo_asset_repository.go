package repository

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "assets"

type MongoAssetRepository struct {
	collection *mongo.Collection
}

func (m MongoAssetRepository) Insert(asset entities.Asset) error {
	_, error := m.collection.InsertOne(nil, asset)
	return error
}

func NewAssetRepository(database mongo.Database) MongoAssetRepository {
	return MongoAssetRepository{
		collection: database.Collection(COLLECTION_NAME),
	}
}
