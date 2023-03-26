package providers

import (
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/FelipeNFL/check-assets-api/adapters"
	"github.com/FelipeNFL/check-assets-api/commom"
	"github.com/FelipeNFL/check-assets-api/domain/protocols"
	"github.com/FelipeNFL/check-assets-api/infra"
)

type HttpAssetInfoProvider struct {
	HttpClient  adapters.HttpClient
	GetInfoFunc func(codes []string) (protocols.AssetInfoResult, error)
}

type YahooFinanceSchema struct {
	QuoteResponse struct {
		Result []struct {
			RegularMarketPrice float64 `json:"regularMarketPrice"`
			Symbol             string  `json:"symbol"`
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

func (p HttpAssetInfoProvider) GetInfo(codes []string) (protocols.AssetInfoResult, error) {
	codesJoined := strings.Join(codes, ",")

	url := "https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=" + codesJoined
	apiKey := commom.GetEnvironmentVariable("YAHOO_FINANCE_API_KEY")
	headers := map[string]string{"X-API-KEY": apiKey}

	body, err := p.HttpClient.Get(url, headers)
	if err != nil {
		log.Error("Error getting price for asset list: ", codesJoined, " - ", err)
		return protocols.AssetInfoResult{}, infra.ErrGetAssetPrice
	}

	parsed := YahooFinanceSchema{}

	json.Unmarshal(body, &parsed)

	assetsInfo := make(map[string]protocols.AssetInfo)

	for _, result := range parsed.QuoteResponse.Result {
		if result.RegularMarketPrice == 0 {
			log.Error("Error getting price for asset: ", result.Symbol, ". Asset code is invalid.")
			return protocols.AssetInfoResult{}, infra.ErrAssetNotFound
		}

		assetsInfo[result.Symbol] = protocols.AssetInfo{
			Price: result.RegularMarketPrice,
		}
	}

	return assetsInfo, nil
}
