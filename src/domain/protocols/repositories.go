package protocols

import "github.com/FelipeNFL/check-assets-api/domain/entities"

type AssetRepository interface {
	Insert(asset entities.Asset) (entities.Asset, error)
	CheckIfAssetExists(code string) (bool, error)
	GetLastPosition() (int, error)
	GetAll() ([]entities.Asset, error)
}

type AssetOrdinationRepository interface {
	Clean() error
	Insert(assetOrdination entities.AssetOrdination) (entities.AssetOrdination, error)
	Get() (entities.AssetOrdination, error)
}
