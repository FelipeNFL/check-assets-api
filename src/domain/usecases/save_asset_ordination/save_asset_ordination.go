package save_asset_ordination

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
)

type SaveAssetOrdinationUseCase struct {
	AssetOrdinationRepository protocols.AssetOrdinationRepository
	SaveFunc                  func(assetOrdination entities.AssetOrdination) (entities.AssetOrdination, error)
}

func (s SaveAssetOrdinationUseCase) Save(assetOrdination string) (entities.AssetOrdination, error) {
	assetOrdinationEntity, err := entities.NewAssetOrdination(
		entities.Ordination(assetOrdination),
	)

	if err != nil {
		return entities.AssetOrdination{}, err
	}

	err = s.AssetOrdinationRepository.Clean()

	if err != nil {
		return entities.AssetOrdination{}, err
	}

	return s.AssetOrdinationRepository.Insert(assetOrdinationEntity)
}

type NewSaveAssetOrdinationUseCaseData struct {
	AssetOrdinationRepository protocols.AssetOrdinationRepository
}

func NewSaveAssetOrdinationUseCase(data NewSaveAssetOrdinationUseCaseData) SaveAssetOrdinationUseCase {
	return SaveAssetOrdinationUseCase{
		AssetOrdinationRepository: data.AssetOrdinationRepository,
	}
}
