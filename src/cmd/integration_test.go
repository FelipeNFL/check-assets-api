package cmd

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FelipeNFL/check-assets-api/commom"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb/asset"
	"github.com/FelipeNFL/check-assets-api/infra/repository/mongodb/asset_ordination"
)

func post(url string, body string) (int, []byte, error) {
	method := "POST"
	payload := strings.NewReader(body)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, result, nil
}

func cleanCollections(database *mongo.Database) {
	database.Collection(asset.COLLECTION_NAME).Drop(context.TODO())
	database.Collection(asset_ordination.COLLECTION_NAME).Drop(context.TODO())
}

func TestIntegration(t *testing.T) {
	URL := "http://localhost:" + os.Getenv("PORT")

	database := commom.GetMongoDatabase("api")

	cleanCollections(database)

	t.Run("POST /asset should return 200", func(t *testing.T) {
		status, body, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)

		type Response struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		var response Response
		json.Unmarshal(body, &response)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)
		assert.Equal(t, response.Code, "GOGL")
		assert.Equal(t, response.Order, 1)

		cleanCollections(database)
	})

	t.Run("POST /asset should return 400 when code is empty", func(t *testing.T) {
		status, body, err := post(
			URL+"/asset",
			`{"code": ""}`,
		)

		type Response struct {
			Message string `json:"error"`
		}

		var response Response
		json.Unmarshal(body, &response)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 400)
		assert.Equal(t, response.Message, "asset code is empty")

		cleanCollections(database)
	})

	t.Run("POST /asset should return 400 when code is invalid", func(t *testing.T) {
		status, body, err := post(
			URL+"/asset",
			`{"code": "ASSET_CODE_INVALID_IN_YAHOO_FINANCE"}`,
		)

		type Response struct {
			Message string `json:"error"`
		}

		var response Response
		json.Unmarshal(body, &response)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 400)
		assert.Equal(t, response.Message, "asset not found")

		cleanCollections(database)
	})

	t.Run("POST /asset should return 400 when code is already registered", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, body, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)

		type Response struct {
			Message string `json:"error"`
		}

		var response Response
		json.Unmarshal(body, &response)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 400)
		assert.Equal(t, response.Message, "asset already created")

		cleanCollections(database)
	})

	t.Run("GET /asset/price with one code should return 200", func(t *testing.T) {
		res, err := http.Get(URL + "/asset/price?code=GOGL")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string  `json:"code"`
			Price float64 `json:"price"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 1)
		assert.Equal(t, response[0].Code, "GOGL")

		cleanCollections(database)
	})

	t.Run("GET /asset/price with many codes should return 200", func(t *testing.T) {
		res, err := http.Get(URL + "/asset/price?code=GOGL,AAPL")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string  `json:"code"`
			Price float64 `json:"price"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 2)
		assert.Equal(t, response[0].Code, "GOGL")
		assert.Equal(t, response[1].Code, "AAPL")

		cleanCollections(database)
	})

	t.Run("GET /asset/price with invalid code should return 400", func(t *testing.T) {
		res, err := http.Get(URL + "/asset/price?code=INVALID_CODE")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 400)

		type Response struct {
			Message string `json:"error"`
		}

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, response.Message, "asset not found")

		cleanCollections(database)
	})

	t.Run("GET /asset/price with empty code should return 400", func(t *testing.T) {
		res, err := http.Get(URL + "/asset/price?code=")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 400)

		type Response struct {
			Message string `json:"error"`
		}

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, response.Message, "asset not found")

		cleanCollections(database)
	})

	t.Run("GET /asset/price valid and invalid code should return only valid", func(t *testing.T) {
		res, err := http.Get(URL + "/asset/price?code=GOGL,INVALID_CODE")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string  `json:"code"`
			Price float64 `json:"price"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 1)
		assert.Equal(t, response[0].Code, "GOGL")

		cleanCollections(database)
	})

	t.Run("GET /asset should return 200", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "AAPL"}`,
		)

		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		res, err := http.Get(URL + "/asset")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 2)

		assert.Equal(t, response[0].Code, "GOGL")
		assert.Equal(t, response[0].Order, 1)

		assert.Equal(t, response[1].Code, "AAPL")
		assert.Equal(t, response[1].Order, 2)

		cleanCollections(database)
	})

	t.Run("GET /asset should return 200 when no asset is registered", func(t *testing.T) {
		res, err := http.Get(URL + "/asset")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 0)

		cleanCollections(database)
	})

	t.Run("POST /asset/ordination with valid alphabetical asc should return 200", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "AAPL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset/ordination",
			`{"ordination": "alphabetical"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		res, err := http.Get(URL + "/asset")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 2)

		assert.Equal(t, response[0].Code, "AAPL")
		assert.Equal(t, response[0].Order, 2)

		assert.Equal(t, response[1].Code, "GOGL")
		assert.Equal(t, response[1].Order, 1)

		cleanCollections(database)
	})

	t.Run("POST /asset/ordination with valid alphabetical desc should return 200", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "AAPL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset/ordination",
			`{"ordination": "alphabetical"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		res, err := http.Get(URL + "/asset?order=desc")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 2)

		assert.Equal(t, response[0].Code, "GOGL")
		assert.Equal(t, response[0].Order, 1)

		assert.Equal(t, response[1].Code, "AAPL")
		assert.Equal(t, response[1].Order, 2)

		cleanCollections(database)
	})

	t.Run("POST /asset/ordination with price asc should return 200", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "AAPL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset/ordination",
			`{"ordination": "price"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		type ItemPrice struct {
			Code  string  `json:"code"`
			Price float64 `json:"price"`
		}

		type ResponsePrice []ItemPrice

		res, err := http.Get(URL + "/asset/price?code=GOGL,AAPL")
		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		var responsePrice ResponsePrice
		json.NewDecoder(res.Body).Decode(&responsePrice)

		var applePosition int
		var googlePosition int

		if responsePrice[0].Code == "GOGL" && responsePrice[0].Price > responsePrice[1].Price {
			applePosition = 0
			googlePosition = 1
		}

		if responsePrice[0].Code == "GOGL" && responsePrice[0].Price < responsePrice[1].Price {
			googlePosition = 0
			applePosition = 1
		}

		res, err = http.Get(URL + "/asset")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 2)

		assert.Equal(t, response[googlePosition].Code, "GOGL")
		assert.Equal(t, response[googlePosition].Order, 1)

		assert.Equal(t, response[applePosition].Code, "AAPL")
		assert.Equal(t, response[applePosition].Order, 2)

		cleanCollections(database)
	})

	t.Run("POST /asset/ordination with price desc should return 200", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "AAPL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset/ordination",
			`{"ordination": "price"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		type ItemPrice struct {
			Code  string  `json:"code"`
			Price float64 `json:"price"`
		}

		type ResponsePrice []ItemPrice

		res, err := http.Get(URL + "/asset/price?code=GOGL,AAPL")
		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		var responsePrice ResponsePrice
		json.NewDecoder(res.Body).Decode(&responsePrice)

		var applePosition int
		var googlePosition int

		if responsePrice[0].Code == "GOGL" && responsePrice[0].Price > responsePrice[1].Price {
			googlePosition = 0
			applePosition = 1
		}

		if responsePrice[0].Code == "GOGL" && responsePrice[0].Price < responsePrice[1].Price {
			applePosition = 0
			googlePosition = 1
		}

		res, err = http.Get(URL + "/asset?order=desc")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 2)

		assert.Equal(t, response[googlePosition].Code, "GOGL")
		assert.Equal(t, response[googlePosition].Order, 1)

		assert.Equal(t, response[applePosition].Code, "AAPL")
		assert.Equal(t, response[applePosition].Order, 2)

		cleanCollections(database)
	})

	t.Run("POST /asset/ordination with custom should return 200", func(t *testing.T) {
		status, _, err := post(
			URL+"/asset",
			`{"code": "GOGL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "AAPL"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset",
			`{"code": "META"}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		status, _, err = post(
			URL+"/asset/ordination",
			`{"ordination": "custom", "custom_order": ["AAPL", "META", "GOGL"]}`,
		)
		assert.Equal(t, err, nil)
		assert.Equal(t, status, 200)

		type Item struct {
			Code  string `json:"code"`
			Order int    `json:"order"`
		}

		type Response []Item

		res, err := http.Get(URL + "/asset")

		assert.Equal(t, err, nil)
		assert.Equal(t, res.StatusCode, 200)

		var response Response
		json.NewDecoder(res.Body).Decode(&response)

		assert.Equal(t, len(response), 3)

		assert.Equal(t, response[0].Code, "AAPL")
		assert.Equal(t, response[0].Order, 0)

		assert.Equal(t, response[1].Code, "META")
		assert.Equal(t, response[1].Order, 1)

		assert.Equal(t, response[2].Code, "GOGL")
		assert.Equal(t, response[2].Order, 2)

		cleanCollections(database)
	})

}
