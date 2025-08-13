package asynq

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/fx"

	"github.com/alfariiizi/vandor/internal/core/job"
	"github.com/alfariiizi/vandor/internal/delivery/worker"
)

// Params for fx
type WorkerParams struct {
	fx.In

	Server    *worker.Server
	Jobs      *job.Jobs
	Lifecycle fx.Lifecycle
}

// WorkerModule starts the worker pool.
// Actual job registrations are auto-generated in worker_gen.go
func RegisterWorker(p WorkerParams) {
	mux := asynq.NewServeMux()

	// Auto-generated registrations:
	registerJobHandlers(mux, p.Jobs)

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() { _ = p.Server.Start(mux) }()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Server.Stop()
			return nil
		},
	})
}

var WorkerModule = fx.Module("asynq",
	fx.Invoke(RegisterWorker),
)
