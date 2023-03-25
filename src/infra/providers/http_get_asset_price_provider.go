package providers

import (
	"encoding/json"
	"log"

	"github.com/FelipeNFL/check-assets-api/adapters"
	"github.com/FelipeNFL/check-assets-api/commom"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/infra"
)

type HttpGetAssetInfoProvider struct {
	HttpClient adapters.HttpClient
}

type YahooFinanceSchema struct {
	QuoteResponse struct {
		Result []struct {
			RegularMarketPrice float64 `json:"regularMarketPrice"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

type NewGetInfoProviderData struct {
	HttpClient adapters.HttpClient
}

func NewGetInfoProvider(data NewGetInfoProviderData) HttpGetAssetInfoProvider {
	return HttpGetAssetInfoProvider{
		HttpClient: data.HttpClient,
	}
}

func (p HttpGetAssetInfoProvider) GetInfo(code string) (protocols.GetInfoProviderResult, error) {
	url := "https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=" + code
	apiKey := commom.GetEnvironmentVariable("YAHOO_FINANCE_API_KEY")
	headers := map[string]string{"X-API-KEY": apiKey}

	body, err := p.HttpClient.Get(url, headers)
	if err != nil {
		log.Fatal("Error getting price for asset: ", code, " - ", err)
		return protocols.GetInfoProviderResult{}, infra.ErrGetAssetPrice{}
	}

	parsed := YahooFinanceSchema{}

	json.Unmarshal(body, &parsed)

	if len(parsed.QuoteResponse.Result) == 0 {
		log.Fatal("Error getting price for asset: ", code, ". Asset is invalid.")
		return protocols.GetInfoProviderResult{}, infra.ErrAssetNotFound{}
	}

	result := protocols.GetInfoProviderResult{
		Price: parsed.QuoteResponse.Result[0].RegularMarketPrice,
	}

	return result, nil
}
