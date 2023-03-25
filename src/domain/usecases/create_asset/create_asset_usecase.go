package create_asset

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/domain/usecases"
)

type CreateAssetUseCase struct {
	AssetRepository   protocols.AssetRepository
	AssetInfoProvider protocols.AssetInfoProvider
	CreateFunc        func(code string) (entities.Asset, error)
}

type NewCreateAssetUseCaseData struct {
	AssetRepository   protocols.AssetRepository
	AssetInfoProvider protocols.AssetInfoProvider
}

func NewCreateAssetUseCase(data NewCreateAssetUseCaseData) CreateAssetUseCase {
	return CreateAssetUseCase{
		AssetRepository:   data.AssetRepository,
		AssetInfoProvider: data.AssetInfoProvider,
	}
}

func (c CreateAssetUseCase) validateAsset(code string) error {
	isAssetAlreadyInserted, checkIfAssetExistsError := c.AssetRepository.CheckIfAssetExists(code)

	if checkIfAssetExistsError != nil {
		return checkIfAssetExistsError
	}

	if isAssetAlreadyInserted {
		return usecases.ErrAssetAlreadyCreated
	}

	_, getAssetInfoProviderError := c.AssetInfoProvider.GetInfo(code)

	if getAssetInfoProviderError != nil {
		return getAssetInfoProviderError
	}

	return nil
}

func (c CreateAssetUseCase) Create(code string) (entities.Asset, error) {
	err := c.validateAsset(code)

	if err != nil {
		return entities.Asset{}, err
	}

	lastPosition, err := c.AssetRepository.GetLastPosition()

	if err != nil {
		return entities.Asset{}, err
	}

	order := lastPosition + 1
	asset, err := entities.NewAsset(code, order)

	if err != nil {
		return entities.Asset{}, err
	}

	return c.AssetRepository.Insert(asset)
}
