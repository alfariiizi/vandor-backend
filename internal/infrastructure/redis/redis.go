package redis

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() *Redis {
	// Setup the client
	cfg := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

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
