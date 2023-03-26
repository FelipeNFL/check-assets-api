package get_asset_list

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/mocks"
)

func TestGetAssetListUsecase(t *testing.T) {
	t.Run("should return a list of assets correctly with default sort", func(t *testing.T) {
		price := 10.0

		assetListInput := []entities.Asset{
			{Code: "code1", Order: 1},
			{Code: "code2", Order: 2},
		}

		assetRepository := mocks.NewMockAssetRepository(mocks.NewMockAssetRepositoryData{
			AssetList: assetListInput,
		})

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{price, price}, nil)

		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:           assetRepository,
			AssetInfoProvider:         getAssetInfoProvider,
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetList, err := getAssetListUseCase.Get("asc")

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 2)
		assert.Equal(t, assetList[0].Order, 1)
		assert.Equal(t, assetList[0].Price, price)
		assert.Equal(t, assetList[1].Order, 2)
		assert.Equal(t, assetList[1].Price, price)
	})

	t.Run("should return a list of assets correctly with alphabetical sort asc", func(t *testing.T) {
		price := 10.0

		assetListInput := []entities.Asset{
			{Code: "b", Order: 1},
			{Code: "a", Order: 2},
		}

		assetRepository := mocks.NewMockAssetRepository(mocks.NewMockAssetRepositoryData{
			AssetList: assetListInput,
		})

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{price, price}, nil)

		assetOrdination, err := entities.NewAssetOrdination("alphabetical", nil)
		assert.Equal(t, err, nil)

		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{AssetOrdination: assetOrdination},
		)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:           assetRepository,
			AssetInfoProvider:         getAssetInfoProvider,
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetList, err := getAssetListUseCase.Get("asc")

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 2)
		assert.Equal(t, assetList[0].Code, "a")
		assert.Equal(t, assetList[0].Order, 2)
		assert.Equal(t, assetList[0].Price, price)
		assert.Equal(t, assetList[1].Code, "b")
		assert.Equal(t, assetList[1].Order, 1)
		assert.Equal(t, assetList[1].Price, price)
	})

	t.Run("should return a list of assets correctly with alphabetical sort desc", func(t *testing.T) {
		price := 10.0

		assetListInput := []entities.Asset{
			{Code: "b", Order: 1},
			{Code: "a", Order: 2},
		}

		assetRepository := mocks.NewMockAssetRepository(mocks.NewMockAssetRepositoryData{
			AssetList: assetListInput,
		})

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{price, price}, nil)

		assetOrdination, err := entities.NewAssetOrdination("alphabetical", nil)
		assert.Equal(t, err, nil)

		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{AssetOrdination: assetOrdination},
		)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:           assetRepository,
			AssetInfoProvider:         getAssetInfoProvider,
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetList, err := getAssetListUseCase.Get("desc")

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 2)
		assert.Equal(t, assetList[0].Code, "b")
		assert.Equal(t, assetList[0].Order, 1)
		assert.Equal(t, assetList[0].Price, price)
		assert.Equal(t, assetList[1].Code, "a")
		assert.Equal(t, assetList[1].Order, 2)
		assert.Equal(t, assetList[1].Price, price)
	})

	t.Run("should return a list of assets correctly with price sort asc", func(t *testing.T) {
		price1 := 10.0
		price2 := 20.0

		assetListInput := []entities.Asset{
			{Code: "code1", Order: 1},
			{Code: "code2", Order: 2},
		}

		assetRepository := mocks.NewMockAssetRepository(mocks.NewMockAssetRepositoryData{
			AssetList: assetListInput,
		})

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{price1, price2}, nil)

		assetOrdination, err := entities.NewAssetOrdination("price", nil)
		assert.Equal(t, err, nil)

		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{AssetOrdination: assetOrdination},
		)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:           assetRepository,
			AssetInfoProvider:         getAssetInfoProvider,
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetList, err := getAssetListUseCase.Get("asc")

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 2)
		assert.Equal(t, assetList[0].Code, "code1")
		assert.Equal(t, assetList[0].Order, 1)
		assert.Equal(t, assetList[0].Price, price1)
		assert.Equal(t, assetList[1].Code, "code2")
		assert.Equal(t, assetList[1].Order, 2)
		assert.Equal(t, assetList[1].Price, price2)
	})

	t.Run("should return a list of assets correctly with price sort desc", func(t *testing.T) {
		price1 := 10.0
		price2 := 20.0

		assetListInput := []entities.Asset{
			{Code: "code1", Order: 1},
			{Code: "code2", Order: 2},
		}

		assetRepository := mocks.NewMockAssetRepository(mocks.NewMockAssetRepositoryData{
			AssetList: assetListInput,
		})

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{price1, price2}, nil)

		assetOrdination, err := entities.NewAssetOrdination("price", nil)
		assert.Equal(t, err, nil)

		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{AssetOrdination: assetOrdination},
		)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:           assetRepository,
			AssetInfoProvider:         getAssetInfoProvider,
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetList, err := getAssetListUseCase.Get("desc")

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 2)
		assert.Equal(t, assetList[0].Code, "code2")
		assert.Equal(t, assetList[0].Order, 2)
		assert.Equal(t, assetList[0].Price, price2)
		assert.Equal(t, assetList[1].Code, "code1")
		assert.Equal(t, assetList[1].Order, 1)
		assert.Equal(t, assetList[1].Price, price1)
	})

	t.Run("should return empty list of assets when asset repository returns empty list", func(t *testing.T) {
		assetRepository := mocks.NewMockAssetRepository(
			mocks.NewMockAssetRepositoryData{AssetList: []entities.Asset{}},
		)

		getAssetInfoProvider := mocks.NewMockAssetInfoProvider(nil, nil)

		assetOrdination, err := entities.NewAssetOrdination("alphabetical", nil)
		assert.Equal(t, err, nil)

		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{AssetOrdination: assetOrdination},
		)

		getAssetListUseCase := NewGetAssetListUseCase(NewGetAssetListUseCaseData{
			AssetRepository:           assetRepository,
			AssetInfoProvider:         getAssetInfoProvider,
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetList, err := getAssetListUseCase.Get("asc")

		assert.Equal(t, err, nil)
		assert.Equal(t, len(assetList), 0)
	})
}
