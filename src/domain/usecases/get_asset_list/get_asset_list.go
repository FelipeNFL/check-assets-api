package get_asset_list

import (
	"log"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	providers "github.com/FelipeNFL/check-assets-api/domain/protocols"
	repositories "github.com/FelipeNFL/check-assets-api/domain/protocols"
)

type GetAssetListUseCase struct {
	AssetRepository      repositories.AssetRepository
	GetAssetInfoProvider providers.GetAssetInfoProvider
	GetFunc              func() ([]entities.Asset, error)
}

type NewGetAssetListUseCaseData struct {
	AssetRepository      repositories.AssetRepository
	GetAssetInfoProvider providers.GetAssetInfoProvider
}

func NewGetAssetListUseCase(data NewGetAssetListUseCaseData) GetAssetListUseCase {
	return GetAssetListUseCase{
		AssetRepository:      data.AssetRepository,
		GetAssetInfoProvider: data.GetAssetInfoProvider,
	}
}

func (g GetAssetListUseCase) Get() ([]entities.Asset, error) {
	assets, err := g.AssetRepository.GetAll()

	if err != nil {
		return nil, err
	}

	for i, asset := range assets {
		assetInfo, err := g.GetAssetInfoProvider.GetInfo(asset.Code)

		if err != nil {
			log.Fatal("Error getting price for asset: ", asset.Code, " - ", err)
			assets[i].Price = 0
		}

		assets[i].Price = assetInfo.Price
	}

	return assets, nil
}
