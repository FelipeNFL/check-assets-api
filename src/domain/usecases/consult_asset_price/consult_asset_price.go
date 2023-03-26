package consult_asset_price

import "github.com/FelipeNFL/check-assets-api/domain/protocols"

type ConsultAssetPriceUseCase struct {
	AssetInfoProvider protocols.AssetInfoProvider
	GetFunc           func(code string) (float64, error)
}

type NewConsultAssetPriceUseCaseData struct {
	AssetInfoProvider protocols.AssetInfoProvider
}

type AssetInfo struct {
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

func (c ConsultAssetPriceUseCase) Get(codes []string) ([]AssetInfo, error) {
	assetInfo, err := c.AssetInfoProvider.GetInfo(codes)

	if err != nil {
		return nil, err
	}

	assetInfoList := make([]AssetInfo, 0)

	for _, code := range codes {
		price := assetInfo[code].Price

		if price != 0 {
			assetInfoList = append(assetInfoList, AssetInfo{
				Code:  code,
				Price: assetInfo[code].Price,
			})
		}
	}

	return assetInfoList, nil
}

func NewConsultAssetPriceUseCase(data NewConsultAssetPriceUseCaseData) ConsultAssetPriceUseCase {
	return ConsultAssetPriceUseCase{
		AssetInfoProvider: data.AssetInfoProvider,
	}
}
