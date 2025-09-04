package model

import (
	"context"

	"github.com/alfariiizi/vandor/internal/types"
	"github.com/hibiken/asynq"
)

type Job[I any] interface {
	// Key returns the unique key for the job, used for routing and identification.
	Key() string
	// Handle processes the job with the given payload.
	Handle(ctx context.Context, payload I) error
	// Enqueue adds the job to the queue with the given payload.
	Enqueue(ctx context.Context, payload I) (*asynq.TaskInfo, error)
	// HTTPRegisterRoute registers the HTTP route for the job.
	HTTPRegisterRoute()
}

type JobHTTPHandlerData struct {
	TaskID string `json:"task_id"`
}

type JobHTTPHandlerResponse types.OutputResponseData[JobHTTPHandlerData]
