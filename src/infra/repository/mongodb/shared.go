package mongodb

import (
	"context"
	"time"
)

const DEFAULT_TIMEOUT = 10

func GetContext(timeout_optional ...int) (context.Context, context.CancelFunc) {
	timeout := DEFAULT_TIMEOUT

	if len(timeout_optional) > 0 {
		timeout = timeout_optional[0]
	}

	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}
