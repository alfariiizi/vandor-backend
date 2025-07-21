package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/alfariiizi/vandor/config"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() *Redis {
	// Setup the client
	cfg := config.GetConfig()
	rdb := redis.NewClient(&cfg.Redis)

	// Test the connection
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis:", pong)
	return &Redis{
		Client: rdb,
	}
}
