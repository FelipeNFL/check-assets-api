package save_asset_ordination

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/domain/usecases"
)

type SaveAssetOrdinationUseCase struct {
	AssetRepository           protocols.AssetRepository
	AssetOrdinationRepository protocols.AssetOrdinationRepository
	SaveFunc                  func(assetOrdination entities.AssetOrdination) (entities.AssetOrdination, error)
}

func (s SaveAssetOrdinationUseCase) storageInDatabase(data entities.AssetOrdination, currentAssets []entities.Asset) (entities.AssetOrdination, error) {
	if err := s.AssetOrdinationRepository.Clean(); err != nil {
		return entities.AssetOrdination{}, err
	}

	if data.Ordination != entities.Custom {
		return s.AssetOrdinationRepository.Insert(data)
	}

	for newOrder, asset := range currentAssets {
		if err := s.AssetRepository.UpdateAssetOrder(asset.Code, newOrder); err != nil {
			return data, err
		}
	}

	return s.AssetOrdinationRepository.Insert(data)
}

func (s SaveAssetOrdinationUseCase) validateCustomOrdenation(
	assetOrdination entities.AssetOrdination,
	currentAssets []entities.Asset,
) error {
	if assetOrdination.Ordination != entities.Custom {
		return nil
	}

	if len(currentAssets) != len(assetOrdination.CustomOrder) {
		return usecases.ErrAssetListHasInvalidSize
	}

	currentAssetMap := make(map[string]struct{}, len(currentAssets))

	for _, asset := range currentAssets {
		if _, repeated := currentAssetMap[asset.Code]; repeated {
			panic("asset repeated in database")
		} else {
			currentAssetMap[asset.Code] = struct{}{}
		}
	}

	customAssetsMap := make(map[string]struct{}, len(assetOrdination.CustomOrder))

	for _, asset := range assetOrdination.CustomOrder {
		if _, repeated := customAssetsMap[asset]; repeated {
			return usecases.ErrThereIsAssetRepetition
		}

		if _, assetFoundInDatabase := currentAssetMap[asset]; !assetFoundInDatabase {
			return usecases.ErrAssetDoesntExistInDatabase
		}

		customAssetsMap[asset] = struct{}{}
	}

	return nil
}

func (s SaveAssetOrdinationUseCase) Save(ordination string, customOrder []string) (entities.AssetOrdination, error) {
	assetOrdination, err := entities.NewAssetOrdination(ordination, customOrder)

	if err != nil {
		return entities.AssetOrdination{}, err
	}

	if assetOrdination.Ordination != entities.Custom {
		return s.storageInDatabase(assetOrdination, nil)
	}

	currentAssets, err := s.AssetRepository.GetAll()

	if err != nil {
		return entities.AssetOrdination{}, err
	}

	if err = s.validateCustomOrdenation(assetOrdination, currentAssets); err != nil {
		return entities.AssetOrdination{}, err
	}

	return s.storageInDatabase(assetOrdination, currentAssets)
}

type NewSaveAssetOrdinationUseCaseData struct {
	AssetRepository           protocols.AssetRepository
	AssetOrdinationRepository protocols.AssetOrdinationRepository
}

func NewSaveAssetOrdinationUseCase(data NewSaveAssetOrdinationUseCaseData) SaveAssetOrdinationUseCase {
	return SaveAssetOrdinationUseCase{
		AssetRepository:           data.AssetRepository,
		AssetOrdinationRepository: data.AssetOrdinationRepository,
	}
}
