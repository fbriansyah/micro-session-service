package redisclient

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrorNotFound = errors.New("data not found in cache")

type RedisAdapter struct {
	db *redis.Client
}

func NewRedisAdapter(rdb *redis.Client) *RedisAdapter {
	return &RedisAdapter{
		db: rdb,
	}
}

func (a *RedisAdapter) SetData(ctx context.Context, key string, data string, duration time.Duration) error {
	return a.db.Set(ctx, key, data, duration).Err()
}

func (a *RedisAdapter) GetData(ctx context.Context, key string) (string, error) {

	strData, err := a.db.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrorNotFound
		}

		return "", err
	}

	return strData, nil
}
