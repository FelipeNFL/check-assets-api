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

func (s SaveAssetOrdinationUseCase) storageInDatabase(data entities.AssetOrdination) (entities.AssetOrdination, error) {
	if err := s.AssetOrdinationRepository.Clean(); err != nil {
		return entities.AssetOrdination{}, err
	}

	return s.AssetOrdinationRepository.Insert(data)
}

func (s SaveAssetOrdinationUseCase) saveCustomOrder(assetOrdination entities.AssetOrdination) error {
	currentAssets, err := s.AssetRepository.GetAll()
	if err != nil {
		return err
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

	customAssetsMap := make(map[string]int, len(assetOrdination.CustomOrder))

	for order, asset := range assetOrdination.CustomOrder {
		if _, repeated := customAssetsMap[asset]; repeated {
			return usecases.ErrThereIsAssetRepetition
		}

		if _, assetFoundInDatabase := currentAssetMap[asset]; !assetFoundInDatabase {
			return usecases.ErrAssetDoesntExistInDatabase
		}

		customAssetsMap[asset] = order
	}

	for _, asset := range currentAssets {
		newOrder := customAssetsMap[asset.Code]

		if err := s.AssetRepository.UpdateAssetOrder(asset.Code, newOrder); err != nil {
			return err
		}
	}

	return nil
}

func (s SaveAssetOrdinationUseCase) Save(ordination string, customOrder []string) (entities.AssetOrdination, error) {
	assetOrdination, err := entities.NewAssetOrdination(ordination, customOrder)

	if err != nil {
		return entities.AssetOrdination{}, err
	}

	if assetOrdination.Ordination == entities.Custom {
		if err := s.saveCustomOrder(assetOrdination); err != nil {
			return entities.AssetOrdination{}, err
		}
	}

	return s.storageInDatabase(assetOrdination)
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
