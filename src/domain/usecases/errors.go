package usecases

import "errors"

var ErrAssetAlreadyCreated = errors.New("asset already created")
var ErrAssetInvalid = errors.New("asset invalid")

var ErrOrdenationTypeInvalid = errors.New("ordenation type is invalid")
var ErrThereIsAssetRepetition = errors.New("there is asset repetition")
var ErrAssetListHasInvalidSize = errors.New("there is asset missing/extra in custom order")
var ErrAssetDoesntExistInDatabase = errors.New("some asset doesnt exist in database")
