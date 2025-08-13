package config

import (
	"github.com/redis/go-redis/v9"
)

func parseRedisURL(
	host string,
	port string,
	username string,
	password string,
	db int,
) (*redis.Options, error) {
	return &redis.Options{
		Addr:     host + ":" + port,
		Username: username,
		Password: password,
		DB:       db,
	}, nil
}
