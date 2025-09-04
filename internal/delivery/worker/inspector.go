package worker

import (
	"github.com/alfariiizi/vandor/internal/config"
	"github.com/hibiken/asynq"
)

type Inspector struct {
	*asynq.Inspector
}

// Provide *asynq.Client for enqueuing jobs
func NewWorkerInspector() *Inspector {
	cfg := config.GetConfig()
	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &Inspector{inspector}
}
