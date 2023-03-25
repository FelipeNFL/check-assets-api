package consult_asset_price

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FelipeNFL/check-assets-api/domain/mocks"
)

func TestNewConsultAssetPrice(t *testing.T) {
	t.Run("should return a new instance of ConsultAssetPriceUseCase", func(t *testing.T) {
		assetInfoProvider := mocks.NewMockAssetInfoProvider(10.0)
		consultAssetPriceUseCase := NewConsultAssetPriceUseCase(
			NewConsultAssetPriceUseCaseData{
				GetAssetInfoProvider: assetInfoProvider,
			},
		)

		assert.NotNil(t, consultAssetPriceUseCase)
	})

	t.Run("should get the asset price correctly", func(t *testing.T) {
		price := 10.0
		assetInfoProvider := mocks.NewMockAssetInfoProvider(price)
		consultAssetPriceUseCase := NewConsultAssetPriceUseCase(
			NewConsultAssetPriceUseCaseData{
				GetAssetInfoProvider: assetInfoProvider,
			},
		)

		assetPrice, err := consultAssetPriceUseCase.Get("code")

		assert.Nil(t, err)
		assert.Equal(t, assetPrice, price)
	})
}
