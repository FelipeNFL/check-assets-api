package protocols

type AssetInfoResult struct {
	Price float64
}

type AssetInfoProvider interface {
	GetInfo(code string) (AssetInfoResult, error)
}
