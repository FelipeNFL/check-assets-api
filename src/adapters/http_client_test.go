package adapters

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/assert/v2"
)

type HttpBinSchema struct {
	Headers map[string]string `json:"headers"`
}

func TestAdapters(t *testing.T) {
	t.Run("test get http client", func(t *testing.T) {
		client := NewHttpClient()
		body, err := client.Get("https://httpbin.org/get", Headers{"test": "test"})
		assert.Equal(t, err, nil)

		var httpBinSchema HttpBinSchema
		json.Unmarshal(body, &httpBinSchema)

		assert.Equal(t, httpBinSchema.Headers["Test"], "test")
	})
}
