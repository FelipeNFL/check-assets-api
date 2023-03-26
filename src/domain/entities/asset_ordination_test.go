package entities

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAssetOrdination(t *testing.T) {
	t.Run("should create a new asset ordination correctly", func(t *testing.T) {
		assetOrdination, err := NewAssetOrdination("alphabetical", nil)
		assert.Equal(t, Alphabetical, assetOrdination.Ordination)
		assert.Equal(t, err, nil)

		assetOrdination, err = NewAssetOrdination("price", nil)
		assert.Equal(t, Price, assetOrdination.Ordination)
		assert.Equal(t, err, nil)

		assetOrdination, err = NewAssetOrdination("custom", []string{"code1", "code2"})
		assert.Equal(t, Custom, assetOrdination.Ordination)
		assert.Equal(t, err, nil)
	})

	t.Run("should return an error when ordination is invalid", func(t *testing.T) {
		assetOrdination, err := NewAssetOrdination("invalid", nil)
		assert.Equal(t, AssetOrdination{}, assetOrdination)
		assert.Equal(t, err, ErrAssetOrdinationInvalid)

		_, err = NewAssetOrdination("custom", nil)
		assert.Equal(t, err, ErrAssetOrdinationInvalid)
	})
}
