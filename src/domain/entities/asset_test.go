package entities

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAsset(t *testing.T) {
	userID := 123
	order := 1
	asset := NewAsset("code", order, userID)
	assert.Equal(t, "code", asset.Code)
	assert.Equal(t, 1, asset.Order)
}
