package mocks

import "github.com/FelipeNFL/check-assets-api/domain/protocols"

type MockAssetInfoProvider struct {
	GetInfoFunc func(codes []string) (protocols.AssetInfoResult, error)
}

func (m MockAssetInfoProvider) GetInfo(codes []string) (protocols.AssetInfoResult, error) {
	return m.GetInfoFunc(codes)
}

func NewMockAssetInfoProvider(price []float64) protocols.AssetInfoProvider {
	return &MockAssetInfoProvider{
		GetInfoFunc: func(codes []string) (protocols.AssetInfoResult, error) {
			assetInfoResult := make(map[string]protocols.AssetInfo, len(codes))

			for i, code := range codes {
				assetInfoResult[code] = protocols.AssetInfo{Price: price[i]}
			}

			return assetInfoResult, nil
		},
	}
}
