package entities

type Asset struct {
	Code  string `json:"code"`
	Order int    `json:"order"`
}

func NewAsset(code string, order int) Asset {
	return Asset{Code: code, Order: order}
}
