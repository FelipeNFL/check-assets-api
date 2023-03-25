package get_asset_list

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/mocks"
)

func TestGetAssetListUsecase(t *testing.T) {
	t.Run("should return a list of assets correctly", func(t *testing.T) {
		price := 10.0

		assetListInput := []entities.Asset{
			{Code: "code1", Order: 1},
			{Code: "code2", Order: 2},
		}

		assetRepository := mocks.NewMockAssetRepository(mocks.NewMockAssetRepositoryData{
			AssetList: assetListInput,
		})

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider(price)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:   assetRepository,
			AssetInfoProvider: getAssetInfoProvider,
		})

		assetList, err := getAssetListUseCase.Get()

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 2)
		assert.Equal(t, assetList[0].Order, 1)
		assert.Equal(t, assetList[0].Price, price)
		assert.Equal(t, assetList[1].Order, 2)
		assert.Equal(t, assetList[1].Price, price)
	})

}
