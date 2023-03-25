package mocks

import "github.com/FelipeNFL/check-assets-api/domain/protocols"

type MockGetAssetInfoProvider struct {
	GetInfoFunc func(code string) (protocols.GetInfoProviderResult, error)
}

func (m MockGetAssetInfoProvider) GetInfo(code string) (protocols.GetInfoProviderResult, error) {
	return m.GetInfoFunc(code)
}

func NewMockGetAssetInfoProvider(price float64) protocols.GetAssetInfoProvider {
	return &MockGetAssetInfoProvider{
		GetInfoFunc: func(code string) (protocols.GetInfoProviderResult, error) {
			return protocols.GetInfoProviderResult{Price: price}, nil
		},
	}
}
