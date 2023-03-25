package providers

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/FelipeNFL/check-assets-api/adapters"
)

func TestGetInfo(t *testing.T) {
	t.Run("should return asset price", func(t *testing.T) {
		assetCode := "PETR4"
		expectedPrice := 25.0
		httpClient := adapters.NewMockHttpClient(
			adapters.NewMockHttpClientData{
				Price: expectedPrice,
			},
		)

		httpGetAssetInfoProviderData := NewAssetInfoData{
			HttpClient: httpClient,
		}
		httpGetAssetInfoProvider := NewAssetInfo(httpGetAssetInfoProviderData)

		price, err := httpGetAssetInfoProvider.GetInfo(assetCode)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		assert.Equal(t, expectedPrice, price)
	})
}
