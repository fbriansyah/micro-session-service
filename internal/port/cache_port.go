package port

import (
	"context"
	"time"
)

type CacheAdapterPort interface {
	// SetData save data to cache
	SetData(ctx context.Context, key string, data string, duration time.Duration) error
	// GetData get saved data from cache
	GetData(ctx context.Context, key string) (string, error)
	// DeleteData delete saved data in cache
	DeleteData(ctx context.Context, key string) (string, error)
}
