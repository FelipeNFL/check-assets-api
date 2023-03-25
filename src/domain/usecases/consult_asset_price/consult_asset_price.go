package consult_asset_price

import "github.com/FelipeNFL/check-assets-api/domain/protocols"

type ConsultAssetPriceUseCase struct {
	GetAssetInfoProvider protocols.GetAssetInfoProvider
	GetFunc              func(code string) (float64, error)
}

type NewConsultAssetPriceUseCaseData struct {
	GetAssetInfoProvider protocols.GetAssetInfoProvider
}

func (c ConsultAssetPriceUseCase) Get(code string) (float64, error) {
	assetInfo, err := c.GetAssetInfoProvider.GetInfo(code)

	if err != nil {
		return 0, err
	}

	return assetInfo.Price, nil
}

func NewConsultAssetPriceUseCase(data NewConsultAssetPriceUseCaseData) ConsultAssetPriceUseCase {
	return ConsultAssetPriceUseCase{
		GetAssetInfoProvider: data.GetAssetInfoProvider,
	}
}
