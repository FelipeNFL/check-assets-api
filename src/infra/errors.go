package infra

type ErrGetAssetPrice struct {
	Err error
}

func (e ErrGetAssetPrice) Error() string {
	return "error to get asset price"
}

type ErrAssetNotFound struct {
	Err error
}

func (e ErrAssetNotFound) Error() string {
	return "asset not found"
}
