package entities

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAsset(t *testing.T) {
	t.Run("should create a new asset correctly", func(t *testing.T) {
		order := 1
		asset := NewAsset("code", order)
		assert.Equal(t, "code", asset.Code)
		assert.Equal(t, 1, asset.Order)
	})
}
