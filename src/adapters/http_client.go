package adapters

import (
	"io"
	"net/http"
)

type Headers map[string]string

type HttpClient struct {
	Get func(url string, headers Headers) ([]byte, error)
}

func NewHttpClient() HttpClient {
	return HttpClient{
		Get: func(url string, headers Headers) ([]byte, error) {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return []byte{}, err
			}

			for key, value := range headers {
				req.Header.Add(key, value)
			}

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				return []byte{}, err
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return []byte{}, err
			}

			return body, err
		},
	}
}
