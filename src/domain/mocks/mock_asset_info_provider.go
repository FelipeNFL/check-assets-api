package mocks

import "github.com/FelipeNFL/check-assets-api/domain/protocols"

type MockAssetInfoProvider struct {
	GetInfoFunc func(code string) (protocols.AssetInfoResult, error)
}

func (m MockAssetInfoProvider) GetInfo(code string) (protocols.AssetInfoResult, error) {
	return m.GetInfoFunc(code)
}

func NewMockAssetInfoProvider(price float64) protocols.GetAssetInfoProvider {
	return &MockAssetInfoProvider{
		GetInfoFunc: func(code string) (protocols.AssetInfoResult, error) {
			return protocols.AssetInfoResult{Price: price}, nil
		},
	}
}
