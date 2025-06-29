package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/alfariiizi/go-service/config"
	"github.com/alfariiizi/go-service/internal/core/repository"
)

func NewDB() *repository.Client {
	cfg := config.GetConfig()

	client, err := repository.Open(cfg.DB.Driver, cfg.DB.URL)
	if err != nil {
		panic(err)
	}

	return client
}
