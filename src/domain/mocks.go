package mocks

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

type MockAssetRepository struct {
	InsertFunc func(asset entities.Asset) (entities.Asset, error)
}

func (m MockAssetRepository) Insert(asset entities.Asset) (entities.Asset, error) {
	return m.InsertFunc(asset)
}

func (m MockAssetRepository) GetLastPosition() (int, error) {
	return 0, nil
}

func (m MockAssetRepository) CheckIfAssetExists(code string) (bool, error) {
	return false, nil
}

func NewMockAssetRepository() MockAssetRepository {
	return MockAssetRepository{
		InsertFunc: func(asset entities.Asset) (entities.Asset, error) {
			return asset, nil
		},
	}
}
