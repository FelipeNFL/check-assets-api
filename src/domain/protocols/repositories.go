package repositories

import "github.com/FelipeNFL/check-assets-api/domain/entities"

type AssetRepository interface {
	Insert(asset entities.Asset) (entities.Asset, error)
	CheckIfAssetExists(code string) (bool, error)
	GetLastPosition() (int, error)
}
