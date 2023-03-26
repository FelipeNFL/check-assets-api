package get_asset_list

import (
	"log"
	"sort"
	"strings"

	"github.com/FelipeNFL/check-assets-api/domain/entities"
	providers "github.com/FelipeNFL/check-assets-api/domain/protocols"
	repositories "github.com/FelipeNFL/check-assets-api/domain/protocols"
)

type GetAssetListUseCase struct {
	AssetRepository           repositories.AssetRepository
	AssetOrdinationRepository repositories.AssetOrdinationRepository
	AssetInfoProvider         providers.AssetInfoProvider
	GetFunc                   func() ([]entities.Asset, error)
}

type NewGetAssetListUseCaseData struct {
	AssetRepository           repositories.AssetRepository
	AssetOrdinationRepository repositories.AssetOrdinationRepository
	AssetInfoProvider         providers.AssetInfoProvider
}

func NewGetAssetListUseCase(data NewGetAssetListUseCaseData) GetAssetListUseCase {
	return GetAssetListUseCase{
		AssetRepository:           data.AssetRepository,
		AssetOrdinationRepository: data.AssetOrdinationRepository,
		AssetInfoProvider:         data.AssetInfoProvider,
	}
}

func (g GetAssetListUseCase) Get(direction string) ([]entities.Asset, error) {
	directionEntity := entities.Asc

	if direction != "" {
		directionEntity = entities.Direction(direction)
	}

	assets, err := g.AssetRepository.GetAll()

	if err != nil {
		return nil, err
	}

	assetsCode := make([]string, len(assets))

	for i, asset := range assets {
		assetsCode[i] = asset.Code
	}

	assetsInfo, err := g.AssetInfoProvider.GetInfo(assetsCode)
	if err != nil {
		log.Fatal("Error getting price for asset list: ", strings.Join(assetsCode, ","), " - ", err)
		return nil, err
	}

	for i, asset := range assets {
		assets[i].Price = assetsInfo[asset.Code].Price
	}

	assetOrdination, err := g.AssetOrdinationRepository.Get()

	if err != nil {
		return nil, err
	}

	sort.Slice(assets, func(i, j int) bool {
		if assetOrdination.Ordination == entities.Alphabetical && directionEntity == entities.Asc {
			return assets[i].Code < assets[j].Code
		}

		if assetOrdination.Ordination == entities.Alphabetical && directionEntity == entities.Desc {
			return assets[i].Code > assets[j].Code
		}

		if assetOrdination.Ordination == entities.Price && directionEntity == entities.Asc {
			return assets[i].Price < assets[j].Price
		}

		if assetOrdination.Ordination == entities.Price && directionEntity == entities.Desc {
			return assets[i].Price > assets[j].Price
		}

		return assets[i].Order < assets[j].Order
	})

	return assets, nil
}
