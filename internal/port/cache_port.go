package port

import (
	"context"
	"time"
)

type CacheAdapterPort interface {
	SetData(ctx context.Context, key string, data string, duration time.Duration) error
	GetData(ctx context.Context, key string) (string, error)
}
