package asset_ordination

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/commom"
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

func TestMongoAssetOrdinationRepository(t *testing.T) {
	database := commom.GetMongoDatabase("test")
	assetOrdinationRepository := NewAssetOrdinationRepository(*database)
	t.Run("test insert asset ordination", func(t *testing.T) {
		assetOrdination := entities.AssetOrdination{
			Ordination: "alphabetical",
		}

		assetOrdinationInserted, err := assetOrdinationRepository.Insert(assetOrdination)
		assert.Equal(t, err, nil)
		assert.Equal(t, assetOrdinationInserted.Ordination, assetOrdination.Ordination)

		cursor, err := assetOrdinationRepository.collection.Find(context.TODO(), map[string]interface{}{
			"ordination": assetOrdination.Ordination,
		})

		assert.Equal(t, err, nil)

		var assetOrdinationFound entities.AssetOrdination
		count := 0

		for cursor.Next(context.TODO()) {
			cursor.Decode(&assetOrdinationFound)
			count++
		}

		assert.Equal(t, count, 1)
		assert.Equal(t, assetOrdinationFound.Ordination, assetOrdination.Ordination)

		assetOrdinationRepository.collection.Drop(context.TODO())
	})

	t.Run("test get asset ordination", func(t *testing.T) {
		assetOrdination := entities.AssetOrdination{
			Ordination: "alphabetical",
		}

		assetOrdinationInserted, err := assetOrdinationRepository.Insert(assetOrdination)
		assert.Equal(t, err, nil)
		assert.Equal(t, assetOrdinationInserted.Ordination, assetOrdination.Ordination)

		assetOrdinationFound, err := assetOrdinationRepository.Get()
		assert.Equal(t, err, nil)
		assert.Equal(t, assetOrdinationFound.Ordination, assetOrdination.Ordination)

		assetOrdinationRepository.collection.Drop(context.TODO())
	})

	t.Run("test clean asset ordination", func(t *testing.T) {
		assetOrdination := entities.AssetOrdination{
			Ordination: "alphabetical",
		}

		assetOrdinationInserted, err := assetOrdinationRepository.Insert(assetOrdination)
		assert.Equal(t, err, nil)
		assert.Equal(t, assetOrdinationInserted.Ordination, assetOrdination.Ordination)

		assetOrdinationRepository.Clean()

		cursor := assetOrdinationRepository.collection.FindOne(context.TODO(), map[string]interface{}{
			"ordination": assetOrdination.Ordination,
		})

		var assetOrdinationFound entities.AssetOrdination
		err = cursor.Decode(&assetOrdinationFound)
		assert.Equal(t, err.Error(), "mongo: no documents in result")

		assetOrdinationRepository.collection.Drop(context.TODO())
	})
}
