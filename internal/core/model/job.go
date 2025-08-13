package model

import (
	"context"

	"github.com/hibiken/asynq"
)

type Job[I any] interface {
	Handle(ctx context.Context, payload I) error
	Enqueue(ctx context.Context, payload I) (*asynq.TaskInfo, error)
}
