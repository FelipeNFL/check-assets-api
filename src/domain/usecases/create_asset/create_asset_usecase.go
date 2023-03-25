package create_asset

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/domain/usecases"
)

type CreateAssetUseCase struct {
	AssetRepository      protocols.AssetRepository
	GetAssetInfoProvider protocols.GetAssetInfoProvider
	CreateFunc           func(code string) (entities.Asset, error)
}

type NewCreateAssetUseCaseData struct {
	AssetRepository      protocols.AssetRepository
	GetAssetInfoProvider protocols.GetAssetInfoProvider
}

func NewCreateAssetUseCase(data NewCreateAssetUseCaseData) CreateAssetUseCase {
	return CreateAssetUseCase{
		AssetRepository:      data.AssetRepository,
		GetAssetInfoProvider: data.GetAssetInfoProvider,
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

	_, getAssetInfoProviderError := c.GetAssetInfoProvider.GetInfo(code)

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
	asset := entities.NewAsset(code, order)

	return c.AssetRepository.Insert(asset)
}
