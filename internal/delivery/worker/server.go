package worker

import (
	"github.com/alfariiizi/vandor/config"
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
			Concurrency: cfg.Worker.Concurrency,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)
	return &Server{srv}
}
