package mocks

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

type MockAssetRepository struct {
	InsertFunc             func(asset entities.Asset) (entities.Asset, error)
	GetLastPositionFunc    func() (int, error)
	CheckIfAssetExistsFunc func(code string) (bool, error)
	GetAllFunc             func() ([]entities.Asset, error)
}

func (m MockAssetRepository) Insert(asset entities.Asset) (entities.Asset, error) {
	return m.InsertFunc(asset)
}

func (m MockAssetRepository) GetLastPosition() (int, error) {
	return m.GetLastPositionFunc()
}

func (m MockAssetRepository) CheckIfAssetExists(code string) (bool, error) {
	return m.CheckIfAssetExistsFunc(code)
}

func (m MockAssetRepository) GetAll() ([]entities.Asset, error) {
	return m.GetAllFunc()
}

type NewMockAssetRepositoryData struct {
	LastPosition           int
	IsAssetAlreadyInserted bool
	AssetList              []entities.Asset
}

func NewMockAssetRepository(data NewMockAssetRepositoryData) MockAssetRepository {
	return MockAssetRepository{
		InsertFunc: func(asset entities.Asset) (entities.Asset, error) {
			return asset, nil
		},
		GetLastPositionFunc: func() (int, error) {
			return data.LastPosition, nil
		},
		CheckIfAssetExistsFunc: func(code string) (bool, error) {
			return data.IsAssetAlreadyInserted, nil
		},
		GetAllFunc: func() ([]entities.Asset, error) {
			return data.AssetList, nil
		},
	}
}
