package cmd

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FelipeNFL/check-assets-api/adapters"
	"github.com/FelipeNFL/check-assets-api/domain/usecases/create_asset"
	"github.com/FelipeNFL/check-assets-api/domain/usecases/get_asset_list"
	"github.com/FelipeNFL/check-assets-api/infra/providers"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb"
)

func NewGetAssetListUseCase(database *mongo.Database) get_asset_list.GetAssetListUseCase {
	httpClient := adapters.NewHttpClient()
	getPricesDataProvider := providers.NewGetInfoProviderData{HttpClient: httpClient}
	getPricesProvider := providers.NewGetInfoProvider(getPricesDataProvider)

	repository := mongodb.NewAssetRepository(*database)
	useCaseData := get_asset_list.NewGetAssetListUseCaseData{
		AssetRepository:      repository,
		GetAssetInfoProvider: getPricesProvider,
	}
	usecase := get_asset_list.NewGetAssetListUseCase(useCaseData)

	return usecase
}

func NewCreateAssetUseCase(database *mongo.Database) create_asset.CreateAssetUseCase {
	repository := mongodb.NewAssetRepository(*database)
	useCaseData := create_asset.NewCreateAssetUseCaseData{AssetRepository: repository}
	usecase := create_asset.NewCreateAssetUseCase(useCaseData)

	return usecase
}
