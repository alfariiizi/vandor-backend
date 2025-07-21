package db

import (
	vandorConfig "github.com/alfariiizi/vandor/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func NewDB() *Client {
	cfg := vandorConfig.GetConfig()

	client, err := Open(cfg.DB.Driver, cfg.DB.URL)
	if err != nil {
		panic(err)
	}

	return client
}
