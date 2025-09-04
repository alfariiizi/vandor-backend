package worker

import (
	"github.com/alfariiizi/vandor/internal/config"
	"github.com/hibiken/asynq"
)

type Client struct {
	*asynq.Client
}

// Provide *asynq.Client for enqueuing jobs
func NewWorkerClient() *Client {
	cfg := config.GetConfig()
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &Client{client}
}
