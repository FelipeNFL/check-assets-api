package asset_ordination

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb"
)

const COLLECTION_NAME = "asset_ordination"

type MongoAssetOrdinationRepository struct {
	collection *mongo.Collection
}

func (m MongoAssetOrdinationRepository) Insert(asset entities.AssetOrdination) (entities.AssetOrdination, error) {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	_, err := m.collection.InsertOne(ctx, asset)

	return asset, err
}

func (m MongoAssetOrdinationRepository) Get() (entities.AssetOrdination, error) {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	ordenation := entities.AssetOrdination{}
	m.collection.FindOne(ctx, bson.D{}).Decode(&ordenation)

	return ordenation, nil
}

func (m MongoAssetOrdinationRepository) Clean() error {
	ctx, cancel := mongodb.GetContext()
	defer cancel()

	return m.collection.Drop(ctx)
}

func NewAssetOrdinationRepository(database mongo.Database) MongoAssetOrdinationRepository {
	return MongoAssetOrdinationRepository{
		collection: database.Collection(COLLECTION_NAME),
	}
}
