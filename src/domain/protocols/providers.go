package protocols

type AssetInfoResult struct {
	Price float64
}

type GetAssetInfoProvider interface {
	GetInfo(code string) (AssetInfoResult, error)
}
