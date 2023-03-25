package cmd

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FelipeNFL/check-assets-api/adapters"
	"github.com/FelipeNFL/check-assets-api/domain/usecases/consult_asset_price"
	"github.com/FelipeNFL/check-assets-api/domain/usecases/create_asset"
	"github.com/FelipeNFL/check-assets-api/domain/usecases/get_asset_list"
	"github.com/FelipeNFL/check-assets-api/domain/usecases/save_asset_ordination"
	"github.com/FelipeNFL/check-assets-api/infra/providers"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb/asset"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb/asset_ordination"
)

func getPricesProvider() providers.HttpAssetInfoProvider {
	httpClient := adapters.NewHttpClient()
	getPricesDataProvider := providers.NewAssetInfoData{HttpClient: httpClient}
	return providers.NewAssetInfo(getPricesDataProvider)
}

func NewGetAssetListUseCase(database *mongo.Database) get_asset_list.GetAssetListUseCase {
	repository := asset.NewAssetRepository(*database)
	useCaseData := get_asset_list.NewGetAssetListUseCaseData{
		AssetRepository:   repository,
		AssetInfoProvider: getPricesProvider(),
	}

	return get_asset_list.NewGetAssetListUseCase(useCaseData)
}

func NewCreateAssetUseCase(database *mongo.Database) create_asset.CreateAssetUseCase {
	repository := asset.NewAssetRepository(*database)
	useCaseData := create_asset.NewCreateAssetUseCaseData{
		AssetRepository:   repository,
		AssetInfoProvider: getPricesProvider(),
	}

	return create_asset.NewCreateAssetUseCase(useCaseData)
}

func NewConsultAssetPriceUseCase() consult_asset_price.ConsultAssetPriceUseCase {
	useCaseData := consult_asset_price.NewConsultAssetPriceUseCaseData{
		AssetInfoProvider: getPricesProvider(),
	}

	return consult_asset_price.NewConsultAssetPriceUseCase(useCaseData)
}

func NewSaveAssetOrdinationUseCase(database *mongo.Database) save_asset_ordination.SaveAssetOrdinationUseCase {
	repository := asset_ordination.NewAssetOrdinationRepository(*database)
	useCaseData := save_asset_ordination.NewSaveAssetOrdinationUseCaseData{
		AssetOrdinationRepository: repository,
	}

	return save_asset_ordination.NewSaveAssetOrdinationUseCase(useCaseData)
}
