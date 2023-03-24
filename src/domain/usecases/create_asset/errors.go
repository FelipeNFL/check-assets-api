package usecases

type ErrAssetAlreadyCreated struct{}

func (e ErrAssetAlreadyCreated) Error() string {
	return "asset already created"
}
