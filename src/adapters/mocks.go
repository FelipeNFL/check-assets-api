package adapters

import "fmt"

type MockHttpClient struct {
	Get func(url string, headers Headers) ([]byte, error)
}

type NewMockHttpClientData struct {
	Price float64
}

func NewMockHttpClient(data NewMockHttpClientData) MockHttpClient {
	return MockHttpClient{
		Get: func(url string, headers Headers) ([]byte, error) {
			body := fmt.Sprintf(
				`
					{
						"quoteRespose": {
							"result": [
								{
									"regularMarketPrice": %f
								}
							]
						}
					}
				`,
				data.Price,
			)

			return []byte(body), nil
		},
	}
}
