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

func (c CreateAssetUseCase) Create(code string, userID int) error {
	lastPosition, error := c.AssetRepository.GetLastPosition()

	if error != nil {
		return error
	}

	order := lastPosition + 1
	asset := entities.NewAsset(code, order, userID)
	return c.AssetRepository.Insert(asset)
}
