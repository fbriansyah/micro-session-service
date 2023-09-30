package redisclient

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/fbriansyah/micro-session-service/util"
	"github.com/redis/go-redis/v9"
)

var cacheDuration time.Duration
var testAdapter *RedisAdapter

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	cacheDuration = config.RefreshTokenDuration
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	testAdapter = NewRedisAdapter(rdb)

	os.Exit(m.Run())
}
