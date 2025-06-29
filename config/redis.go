package config

import (
	"net/url"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func parseRedisURL(redisURL string) (*redis.Options, error) {
	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	db := 0
	if len(u.Path) > 1 {
		db, err = strconv.Atoi(u.Path[1:])
		if err != nil {
			return nil, err
		}
	}

	password := ""
	if u.User != nil {
		password, _ = u.User.Password()
	}

	return &redis.Options{
		Addr:     u.Host,
		Password: password,
		DB:       db,
	}, nil
}
