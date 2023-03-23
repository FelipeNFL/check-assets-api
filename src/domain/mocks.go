package mocks

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

type MockNewAssetRepository struct {
	InsertFunc func(asset entities.Asset) error
}

func (m MockNewAssetRepository) Insert(asset entities.Asset) error {
	return m.InsertFunc(asset)
}

func (m MockNewAssetRepository) GetLastPosition() (int, error) {
	return 0, nil
}
