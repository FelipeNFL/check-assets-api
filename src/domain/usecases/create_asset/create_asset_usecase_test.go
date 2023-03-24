package usecases

import (
	"testing"

	mocks "github.com/FelipeNFL/check-assets-api/domain"
	"github.com/go-playground/assert/v2"
)

func TestCreateAssetUseCase(t *testing.T) {
	lastPosition := 10
	assetRepository := mocks.NewMockAssetRepository(lastPosition, false)
	createAssetUseCase := NewCreateAssetUseCase(assetRepository)
	asset, error := createAssetUseCase.Create("code")

	assert.Equal(t, nil, error)
	assert.Equal(t, "code", asset.Code)
	assert.Equal(t, lastPosition+1, asset.Order)
}

func TestCreateAssetUseCaseWithAssetAlreadyCreated(t *testing.T) {
	assetRepository := mocks.NewMockAssetRepository(0, true)
	createAssetUseCase := NewCreateAssetUseCase(assetRepository)
	_, error := createAssetUseCase.Create("code")

	assert.Equal(t, ErrAssetAlreadyCreated{}, error)
}
