package consult_asset_price

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FelipeNFL/check-assets-api/domain/mocks"
)

func TestConsultAssetPriceUsecase(t *testing.T) {
	t.Run("should return a new instance of ConsultAssetPriceUseCase", func(t *testing.T) {
		assetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{10.0}, nil)
		consultAssetPriceUseCase := NewConsultAssetPriceUseCase(
			NewConsultAssetPriceUseCaseData{
				AssetInfoProvider: assetInfoProvider,
			},
		)

		assert.NotNil(t, consultAssetPriceUseCase)
	})

	t.Run("should get the asset price correctly", func(t *testing.T) {
		price := 10.0
		assetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{price}, nil)
		consultAssetPriceUseCase := NewConsultAssetPriceUseCase(
			NewConsultAssetPriceUseCaseData{
				AssetInfoProvider: assetInfoProvider,
			},
		)

		assetPrice, err := consultAssetPriceUseCase.Get([]string{"code"})

		assert.Nil(t, err)
		assert.Equal(t, len(assetPrice), 1)
		assert.Equal(t, assetPrice[0].Price, price)
		assert.Equal(t, assetPrice[0].Code, "code")
	})

	t.Run("should return an error when assetInfoProvider returns an error", func(t *testing.T) {
		errorExpected := errors.New("error")
		assetInfoProvider := mocks.NewMockAssetInfoProvider([]float64{10.0}, errorExpected)
		consultAssetPriceUseCase := NewConsultAssetPriceUseCase(
			NewConsultAssetPriceUseCaseData{
				AssetInfoProvider: assetInfoProvider,
			},
		)

		_, err := consultAssetPriceUseCase.Get([]string{"code"})

		assert.NotNil(t, err)
		assert.Equal(t, err, errorExpected)
	})
}
