package protocols

type GetInfoProviderResult struct {
	Price float64
}

type GetAssetInfoProvider interface {
	GetInfo(code string) (GetInfoProviderResult, error)
}
