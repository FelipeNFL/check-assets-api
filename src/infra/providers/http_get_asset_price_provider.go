package providers

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/FelipeNFL/check-assets-api/adapters"
	"github.com/FelipeNFL/check-assets-api/commom"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/infra"
)

type HttpAssetInfoProvider struct {
	HttpClient adapters.HttpClient
}

type YahooFinanceSchema struct {
	QuoteResponse struct {
		Result []struct {
			RegularMarketPrice float64 `json:"regularMarketPrice"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

type NewAssetInfoData struct {
	HttpClient adapters.HttpClient
}

func NewAssetInfo(data NewAssetInfoData) HttpAssetInfoProvider {
	return HttpAssetInfoProvider{
		HttpClient: data.HttpClient,
	}
}

func (p HttpAssetInfoProvider) GetInfo(code string) (protocols.AssetInfoResult, error) {
	url := "https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=" + code
	apiKey := commom.GetEnvironmentVariable("YAHOO_FINANCE_API_KEY")
	headers := map[string]string{"X-API-KEY": apiKey}

	body, err := p.HttpClient.Get(url, headers)
	if err != nil {
		log.Error("Error getting price for asset: ", code, " - ", err)
		return protocols.AssetInfoResult{}, infra.ErrGetAssetPrice
	}

	parsed := YahooFinanceSchema{}

	json.Unmarshal(body, &parsed)

	if len(parsed.QuoteResponse.Result) == 0 {
		log.Error("Error getting price for asset: ", code, ". Asset code is invalid.")
		return protocols.AssetInfoResult{}, infra.ErrAssetNotFound
	}

	result := protocols.AssetInfoResult{
		Price: parsed.QuoteResponse.Result[0].RegularMarketPrice,
	}

	return result, nil
}
