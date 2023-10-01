package main

import (
	"github.com/fbriansyah/micro-session-service/internal/adapter/grpc"
	redisclient "github.com/fbriansyah/micro-session-service/internal/adapter/redis"
	"github.com/fbriansyah/micro-session-service/internal/adapter/token"
	"github.com/fbriansyah/micro-session-service/internal/application"
	"github.com/fbriansyah/micro-session-service/util"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err.Error())
	}

	// create cache client and adapter
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()
	cacheAdapter := redisclient.NewRedisAdapter(rdb)

	// create token maker adapter
	tokenMakerAdapter, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		panic(err)
	}

	// create session service
	sessionService := application.NewSessionService(tokenMakerAdapter, cacheAdapter)

	// create grpc server
	serverAdapter := grpc.NewGrpcServerAdapter(sessionService, grpc.GrpcServerConfig{
		GrpcPort:            config.GrpcPort,
		AccessTokenDuration: config.AccessTokenDuration,
		RefeshTokenDuration: config.RefreshTokenDuration,
	})

	serverAdapter.Run()
}
