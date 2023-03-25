package usecases

import "errors"

var ErrAssetAlreadyCreated = errors.New("asset already created")
var ErrAssetInvalid = errors.New("asset invalid")
