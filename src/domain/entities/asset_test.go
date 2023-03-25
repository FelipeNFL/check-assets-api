package entities

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAsset(t *testing.T) {
	t.Run("should create a new asset correctly", func(t *testing.T) {
		order := 1
		asset, err := NewAsset("code", order)
		assert.Equal(t, "code", asset.Code)
		assert.Equal(t, 1, asset.Order)
		assert.Equal(t, err, nil)
	})

	t.Run("should return an error when code is invalid", func(t *testing.T) {
		order := 1
		asset, err := NewAsset("", order)
		assert.Equal(t, "", asset.Code)
		assert.Equal(t, 0, asset.Order)
		assert.Equal(t, err, ErrAssetCodeInvalid)
	})

	t.Run("should return an error when order is invalid", func(t *testing.T) {
		order := -1
		asset, err := NewAsset("code", order)
		assert.Equal(t, "", asset.Code)
		assert.Equal(t, 0, asset.Order)
		assert.Equal(t, err, ErrAssetOrderInvalid)
	})
}
