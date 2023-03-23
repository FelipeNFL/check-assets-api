package usecases

import (
	"testing"

	mocks "github.com/FelipeNFL/check-assets-api/domain"
)

func TestCreateAssetUseCase(t *testing.T) {
	userID := 123
	assetRepository := mocks.MockNewAssetRepository{}
	createAssetUseCase := NewCreateAssetUseCase(assetRepository)
	createAssetUseCase.Create("code", userID)
}
