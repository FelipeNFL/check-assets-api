package asset

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb"
)

const COLLECTION_NAME = "asset"

type MongoAssetRepository struct {
	collection *mongo.Collection
}

func (m MongoAssetRepository) Insert(asset entities.Asset) (entities.Asset, error) {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	_, err := m.collection.InsertOne(ctx, asset)

	return asset, err
}

func (m MongoAssetRepository) GetLastPosition() (int, error) {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	lastAssetOfList := entities.Asset{}
	opts := options.FindOne().SetSort(bson.D{{Key: "order", Value: -1}})
	m.collection.FindOne(ctx, bson.D{}, opts).Decode(&lastAssetOfList)

	return lastAssetOfList.Order, nil
}

func (m MongoAssetRepository) CheckIfAssetExists(code string) (bool, error) {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	asset := entities.Asset{}
	m.collection.FindOne(ctx, bson.D{{Key: "code", Value: code}}).Decode(&asset)

	return asset.Code != "", nil
}

func (m MongoAssetRepository) GetAll() ([]entities.Asset, error) {
	ctx, cancel := mongodb.GetContext()
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

func (m MongoAssetRepository) UpdateAssetOrder(code string, newOrder int) error {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	_, err := m.collection.UpdateOne(
		ctx,
		bson.D{{Key: "code", Value: code}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "order", Value: newOrder}}}},
	)

	return err
}

func NewAssetRepository(database mongo.Database) MongoAssetRepository {
	return MongoAssetRepository{
		collection: database.Collection(COLLECTION_NAME),
	}
}
