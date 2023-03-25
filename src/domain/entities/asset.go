package entities

type Asset struct {
	Code  string  `json:"code"`
	Order int     `json:"order"`
	Price float64 `json:"price,omitempty" `
}

func NewAsset(code string, order int) (Asset, error) {
	if order < 0 {
		return Asset{}, ErrAssetOrderInvalid
	}

	if code == "" {
		return Asset{}, ErrAssetCodeInvalid
	}

	return Asset{Code: code, Order: order}, nil
}
