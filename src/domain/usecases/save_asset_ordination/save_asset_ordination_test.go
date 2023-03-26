package save_asset_ordination

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/mocks"
	"github.com/FelipeNFL/check-assets-api/domain/usecases"
)

func TestSaveAssetOrdinationUsecase(t *testing.T) {
	t.Run("should save asset ordination correctly", func(t *testing.T) {
		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(NewSaveAssetOrdinationUseCaseData{
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetOrdination, err := saveAssetOrdinationUseCase.Save("alphabetical", []string{})

		assert.Equal(t, err, nil)
		assert.Equal(t, assetOrdination.Ordination, entities.Alphabetical)
	})

	t.Run("should save asset ordination correctly with custom order", func(t *testing.T) {
		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(NewSaveAssetOrdinationUseCaseData{
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		assetOrdination, err := saveAssetOrdinationUseCase.Save("alphabetical", []string{"a", "b"})

		assert.Equal(t, err, nil)
		assert.Equal(t, assetOrdination.Ordination, entities.Alphabetical)
		assert.Equal(t, assetOrdination.CustomOrder[0], "a")
		assert.Equal(t, assetOrdination.CustomOrder[1], "b")
	})

	t.Run("should return error when ordination is not valid", func(t *testing.T) {
		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(NewSaveAssetOrdinationUseCaseData{
			AssetOrdinationRepository: assetOrdinationRepository,
		})

		_, err := saveAssetOrdinationUseCase.Save("invalid", []string{})

		assert.Equal(t, err, entities.ErrAssetOrdinationInvalid)
	})

	t.Run("should return error when custom order has repeated items", func(t *testing.T) {
		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		assetRepository := mocks.NewMockAssetRepository(
			mocks.NewMockAssetRepositoryData{
				AssetList: []entities.Asset{
					{Code: "a"},
					{Code: "b"},
				},
			},
		)

		saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(NewSaveAssetOrdinationUseCaseData{
			AssetOrdinationRepository: assetOrdinationRepository,
			AssetRepository:           assetRepository,
		})

		_, err := saveAssetOrdinationUseCase.Save("custom", []string{"a", "a"})

		assert.Equal(t, err, usecases.ErrThereIsAssetRepetition)
	})

	t.Run("should return error when custom order has invalid items", func(t *testing.T) {
		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		assetRepository := mocks.NewMockAssetRepository(
			mocks.NewMockAssetRepositoryData{
				AssetList: []entities.Asset{
					{Code: "a"},
					{Code: "b"},
				},
			},
		)

		saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(NewSaveAssetOrdinationUseCaseData{
			AssetOrdinationRepository: assetOrdinationRepository,
			AssetRepository:           assetRepository,
		})

		_, err := saveAssetOrdinationUseCase.Save("custom", []string{"a", "c"})

		assert.Equal(t, err, usecases.ErrAssetDoesntExistInDatabase)
	})

	t.Run("should return error when custom order has missing items", func(t *testing.T) {
		assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
			mocks.NewMockAssetOrdinationRepositoryData{},
		)

		assetRepository := mocks.NewMockAssetRepository(
			mocks.NewMockAssetRepositoryData{
				AssetList: []entities.Asset{
					{Code: "a"},
					{Code: "b"},
				},
			},
		)

		saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(NewSaveAssetOrdinationUseCaseData{
			AssetOrdinationRepository: assetOrdinationRepository,
			AssetRepository:           assetRepository,
		})

		_, err := saveAssetOrdinationUseCase.Save("custom", []string{"a"})

		assert.Equal(t, err, usecases.ErrAssetListHasInvalidSize)
	})

	t.Run("should return panic when database has repeated items", func(t *testing.T) {
		assert.PanicMatches(
			t,
			func() {
				assetOrdinationRepository := mocks.NewMockAssetOrdinationRepository(
					mocks.NewMockAssetOrdinationRepositoryData{
						AssetOrdination: entities.AssetOrdination{
							Ordination: entities.Custom,
							CustomOrder: []string{
								"a",
								"b",
							},
						},
					},
				)

				assetRepository := mocks.NewMockAssetRepository(
					mocks.NewMockAssetRepositoryData{
						AssetList: []entities.Asset{
							{Code: "a"},
							{Code: "a"},
						},
					},
				)

				saveAssetOrdinationUseCase := NewSaveAssetOrdinationUseCase(
					NewSaveAssetOrdinationUseCaseData{
						AssetOrdinationRepository: assetOrdinationRepository,
						AssetRepository:           assetRepository,
					},
				)

				saveAssetOrdinationUseCase.Save("custom", []string{"a", "b"})
			},
			"asset repeated in database",
		)
	})

}
