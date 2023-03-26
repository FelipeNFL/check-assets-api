package asset

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/commom"
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

func TestMongoAssetRepository(t *testing.T) {
	database := commom.GetMongoDatabase("test")
	assetRepository := NewAssetRepository(*database)

	t.Run("test insert asset", func(t *testing.T) {
		asset := entities.Asset{
			Code:  "PETR4",
			Order: 1,
		}

		assetInserted, err := assetRepository.Insert(asset)
		assert.Equal(t, err, nil)

		assert.Equal(t, assetInserted.Code, asset.Code)
		assert.Equal(t, assetInserted.Order, asset.Order)

		cursor, err := assetRepository.collection.Find(context.TODO(), map[string]interface{}{
			"code": asset.Code,
		})

		assert.Equal(t, err, nil)

		var assetFound entities.Asset
		count := 0

		for cursor.Next(context.TODO()) {
			cursor.Decode(&assetFound)
			count++
		}

		assert.Equal(t, count, 1)
		assert.Equal(t, assetFound.Code, asset.Code)
		assert.Equal(t, assetFound.Order, asset.Order)

		assetRepository.collection.Drop(context.TODO())
	})

	t.Run("test get last position", func(t *testing.T) {
		asset := entities.Asset{
			Code:  "PETR4",
			Order: 1,
		}

		assetInserted, err := assetRepository.Insert(asset)
		assert.Equal(t, err, nil)

		assert.Equal(t, assetInserted.Code, asset.Code)
		assert.Equal(t, assetInserted.Order, asset.Order)

		lastPosition, err := assetRepository.GetLastPosition()

		assert.Equal(t, err, nil)
		assert.Equal(t, lastPosition, 1)

		assetRepository.collection.Drop(context.TODO())
	})

	t.Run("test check if asset exists", func(t *testing.T) {
		asset := entities.Asset{
			Code:  "PETR4",
			Order: 1,
		}

		assetInserted, err := assetRepository.Insert(asset)
		assert.Equal(t, err, nil)

		assert.Equal(t, assetInserted.Code, asset.Code)
		assert.Equal(t, assetInserted.Order, asset.Order)

		exists, err := assetRepository.CheckIfAssetExists(asset.Code)

		assert.Equal(t, err, nil)
		assert.Equal(t, exists, true)

		assetRepository.collection.Drop(context.TODO())
	})

	t.Run("test get all assets", func(t *testing.T) {
		asset := entities.Asset{
			Code:  "PETR4",
			Order: 1,
		}

		assetInserted, err := assetRepository.Insert(asset)
		assert.Equal(t, err, nil)

		assert.Equal(t, assetInserted.Code, asset.Code)
		assert.Equal(t, assetInserted.Order, asset.Order)

		assets, err := assetRepository.GetAll()

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assets), 1)
		assert.Equal(t, assets[0].Code, asset.Code)
		assert.Equal(t, assets[0].Order, asset.Order)

		assetRepository.collection.Drop(context.TODO())
	})

	t.Run("test update asset order", func(t *testing.T) {
		asset := entities.Asset{
			Code:  "PETR4",
			Order: 1,
		}

		assetInserted, err := assetRepository.Insert(asset)
		assert.Equal(t, err, nil)

		assert.Equal(t, assetInserted.Code, asset.Code)
		assert.Equal(t, assetInserted.Order, asset.Order)

		err = assetRepository.UpdateAssetOrder(asset.Code, 2)

		assert.Equal(t, err, nil)

		assetUpdated, err := assetRepository.GetAll()

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetUpdated), 1)
		assert.Equal(t, assetUpdated[0].Code, asset.Code)

		assetRepository.collection.Drop(context.TODO())
	})
}
