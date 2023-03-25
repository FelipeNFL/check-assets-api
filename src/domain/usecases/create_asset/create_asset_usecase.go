package create_asset

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	repositories "github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/domain/usecases"
)

type CreateAssetUseCase struct {
	AssetRepository repositories.AssetRepository
	CreateFunc      func(code string) (entities.Asset, error)
}

type NewCreateAssetUseCaseData struct {
	AssetRepository repositories.AssetRepository
}

func NewCreateAssetUseCase(data NewCreateAssetUseCaseData) CreateAssetUseCase {
	return CreateAssetUseCase{AssetRepository: data.AssetRepository}
}

func (c CreateAssetUseCase) Create(code string) (entities.Asset, error) {
	isAssetAlreadyCreated, err := c.AssetRepository.CheckIfAssetExists(code)

	if err != nil {
		return entities.Asset{}, err
	}

	if isAssetAlreadyCreated {
		return entities.Asset{}, usecases.ErrAssetAlreadyCreated{}
	}

	lastPosition, err := c.AssetRepository.GetLastPosition()

	if err != nil {
		return entities.Asset{}, err
	}

	order := lastPosition + 1
	asset := entities.NewAsset(code, order)

	return c.AssetRepository.Insert(asset)
}
