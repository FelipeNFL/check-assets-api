package mocks

import (
	"github.com/FelipeNFL/check-assets-api/domain/entities"
)

type MockAssetOrdinationRepository struct {
	CleanFunc  func() error
	InsertFunc func(assetOrdination entities.AssetOrdination) (entities.AssetOrdination, error)
	GetFunc    func() (entities.AssetOrdination, error)
}

func (m MockAssetOrdinationRepository) Clean() error {
	return m.CleanFunc()
}

func (m MockAssetOrdinationRepository) Insert(assetOrdination entities.AssetOrdination) (entities.AssetOrdination, error) {
	return m.InsertFunc(assetOrdination)
}

func (m MockAssetOrdinationRepository) Get() (entities.AssetOrdination, error) {
	return m.GetFunc()
}

type NewMockAssetOrdinationRepositoryData struct {
	AssetOrdination entities.AssetOrdination
}

func NewMockAssetOrdinationRepository(data NewMockAssetOrdinationRepositoryData) MockAssetOrdinationRepository {
	return MockAssetOrdinationRepository{
		CleanFunc: func() error {
			return nil
		},
		InsertFunc: func(assetOrdination entities.AssetOrdination) (entities.AssetOrdination, error) {
			return assetOrdination, nil
		},
		GetFunc: func() (entities.AssetOrdination, error) {
			return data.AssetOrdination, nil
		},
	}
}
