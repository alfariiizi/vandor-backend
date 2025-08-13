package worker

import (
	"github.com/alfariiizi/vandor/config"
	"github.com/hibiken/asynq"
)

type Client struct {
	*asynq.Client
}

// Provide *asynq.Client for enqueuing jobs
func NewWorkerClient() *Client {
	cfg := config.GetConfig()
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.Redis.Addr,
	})
	return &Client{client}
}
