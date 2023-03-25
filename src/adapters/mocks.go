package adapters

import "fmt"

type MockHttpClient struct {
	GetFunc func(url string, headers Headers) ([]byte, error)
}

func (m MockHttpClient) Get(url string, headers Headers) ([]byte, error) {
	return m.GetFunc(url, headers)
}

type NewMockHttpClientData struct {
	Price float64
}

func NewMockHttpClient(data NewMockHttpClientData) MockHttpClient {
	return MockHttpClient{
		GetFunc: func(url string, headers Headers) ([]byte, error) {
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
