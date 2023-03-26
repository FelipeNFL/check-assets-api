package commom

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCommom(t *testing.T) {
	t.Run("test get mongo database", func(t *testing.T) {
		database := GetMongoDatabase("test")
		assert.Equal(t, database.Name(), "test")
	})

	t.Run("test get environment variable", func(t *testing.T) {
		os.Setenv("TEST", "test")
		assert.Equal(t, GetEnvironmentVariable("TEST"), "test")
	})
}
