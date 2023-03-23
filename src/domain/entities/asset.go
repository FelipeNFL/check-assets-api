package entities

type Asset struct {
	Code   string
	Order  int
	UserID int
}

func NewAsset(code string, order int, userID int) Asset {
	return Asset{Code: code, Order: order, UserID: userID}
}
