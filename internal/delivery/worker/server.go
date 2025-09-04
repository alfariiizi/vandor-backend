package worker

import (
	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/hibiken/asynq"
)

type Server struct {
	*asynq.Server
}

func NewWorkerServer() *Server {
	cfg := config.GetConfig()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.Redis.Addr,
			Username: cfg.Redis.Username,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: cfg.Jobs.Concurrency,
			// Optionally specify multiple queues with different priority.
			Queues: utils.GetAllQueueConfig(),
			// See the godoc for other configuration options
		},
	)
	return &Server{srv}
}
