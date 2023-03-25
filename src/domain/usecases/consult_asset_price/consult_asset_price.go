package consult_asset_price

import "github.com/FelipeNFL/check-assets-api/domain/protocols"

type ConsultAssetPriceUseCase struct {
	AssetInfoProvider protocols.AssetInfoProvider
	GetFunc           func(code string) (float64, error)
}

type NewConsultAssetPriceUseCaseData struct {
	AssetInfoProvider protocols.AssetInfoProvider
}

func (c ConsultAssetPriceUseCase) Get(code string) (float64, error) {
	assetInfo, err := c.AssetInfoProvider.GetInfo(code)

	if err != nil {
		return 0, err
	}

	return assetInfo.Price, nil
}

func NewConsultAssetPriceUseCase(data NewConsultAssetPriceUseCaseData) ConsultAssetPriceUseCase {
	return ConsultAssetPriceUseCase{
		AssetInfoProvider: data.AssetInfoProvider,
	}
}
