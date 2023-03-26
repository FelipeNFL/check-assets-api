package protocols

type AssetInfo struct {
	Price float64
}

type AssetInfoResult map[string]AssetInfo

type AssetInfoProvider interface {
	GetInfo(code []string) (AssetInfoResult, error)
}
