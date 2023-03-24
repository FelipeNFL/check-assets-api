package usecases

import (
	"testing"

	mocks "github.com/FelipeNFL/check-assets-api/domain"
	"github.com/go-playground/assert/v2"
)

func TestCreateAssetUseCase(t *testing.T) {
	assetRepository := mocks.NewMockAssetRepository()
	createAssetUseCase := NewCreateAssetUseCase(assetRepository)
	error := createAssetUseCase.Create("code")
	assert.Equal(t, nil, error)
}
