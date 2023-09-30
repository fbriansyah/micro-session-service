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
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     addr,
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	return &RedisAdapter{
		db: rdb,
	}
}

func (a *RedisAdapter) SetData(key string, data string, duration time.Duration) error {
	return a.db.Set(context.Background(), key, data, duration).Err()
}

func (a *RedisAdapter) GetData(key string) (string, error) {

	strData, err := a.db.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrorNotFound
		}

		return "", err
	}

	return strData, nil
}
