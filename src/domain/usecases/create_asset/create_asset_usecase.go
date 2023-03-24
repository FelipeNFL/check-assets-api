package usecases

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	repositories "github.com/FelipeNFL/check-assets-api/domain/protocols"
)

type CreateAssetUseCase struct {
	AssetRepository repositories.AssetRepository
}

func NewCreateAssetUseCase(assetRepository repositories.AssetRepository) CreateAssetUseCase {
	return CreateAssetUseCase{AssetRepository: assetRepository}
}

func (c CreateAssetUseCase) Create(code string) (entities.Asset, error) {
	lastPosition, err := c.AssetRepository.GetLastPosition()

	if err != nil {
		return entities.Asset{}, err
	}

	order := lastPosition + 1
	asset := entities.NewAsset(code, order)

	return c.AssetRepository.Insert(asset)
}
