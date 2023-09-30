package port

import (
	"time"
)

type CacheAdapterPort interface {
	SetData(key string, data string, duration time.Duration) error
	GetData(key string) (string, error)
}
