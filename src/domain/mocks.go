package mocks

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

type MockAssetRepository struct {
	InsertFunc             func(asset entities.Asset) (entities.Asset, error)
	GetLastPositionFunc    func() (int, error)
	CheckIfAssetExistsFunc func(code string) (bool, error)
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

func NewMockAssetRepository(lastPosition int, isAssetAlreadyInserted bool) MockAssetRepository {
	return MockAssetRepository{
		InsertFunc: func(asset entities.Asset) (entities.Asset, error) {
			return asset, nil
		},
		GetLastPositionFunc: func() (int, error) {
			return lastPosition, nil
		},
		CheckIfAssetExistsFunc: func(code string) (bool, error) {
			return isAssetAlreadyInserted, nil
		},
	}
}
