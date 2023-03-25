package create_asset

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/domain/mocks"
	"github.com/FelipeNFL/check-assets-api/domain/usecases"
)

func TestCreateAssetUseCase(t *testing.T) {
	t.Run("should create a new asset correctly", func(t *testing.T) {
		lastPosition := 10
		assetRepository := mocks.NewMockAssetRepository(
			mocks.NewMockAssetRepositoryData{
				LastPosition:           lastPosition,
				IsAssetAlreadyInserted: false,
			},
		)
		createAssetUseCaseData := NewCreateAssetUseCaseData{
			AssetRepository: assetRepository,
		}

		createAssetUseCase := NewCreateAssetUseCase(createAssetUseCaseData)
		asset, err := createAssetUseCase.Create("code")

		assert.Equal(t, nil, err)
		assert.Equal(t, "code", asset.Code)
		assert.Equal(t, lastPosition+1, asset.Order)
	})

	t.Run("should return an error when asset already exists", func(t *testing.T) {
		assetRepository := mocks.NewMockAssetRepository(
			mocks.NewMockAssetRepositoryData{
				IsAssetAlreadyInserted: true,
			},
		)
		createAssetUseCaseData := NewCreateAssetUseCaseData{
			AssetRepository: assetRepository,
		}

		createAssetUseCase := NewCreateAssetUseCase(createAssetUseCaseData)
		_, err := createAssetUseCase.Create("code")

		assert.Equal(t, usecases.ErrAssetAlreadyCreated, err)
	})
}
