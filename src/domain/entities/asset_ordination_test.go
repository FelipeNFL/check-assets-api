package entities

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAssetOrdination(t *testing.T) {
	t.Run("should create a new asset ordination correctly", func(t *testing.T) {
		assetOrdination, err := NewAssetOrdination(Alphabetical)
		assert.Equal(t, Alphabetical, assetOrdination.Ordination)
		assert.Equal(t, err, nil)

		assetOrdination, err = NewAssetOrdination(Price)
		assert.Equal(t, Price, assetOrdination.Ordination)
		assert.Equal(t, err, nil)

		assetOrdination, err = NewAssetOrdination(Custom)
		assert.Equal(t, Custom, assetOrdination.Ordination)
		assert.Equal(t, err, nil)
	})

	t.Run("should return an error when ordination is invalid", func(t *testing.T) {
		assetOrdination, err := NewAssetOrdination("invalid")
		assert.Equal(t, AssetOrdination{}, assetOrdination)
		assert.Equal(t, err, ErrAssetOrdinationInvalid)
	})
}
