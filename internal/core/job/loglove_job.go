package job

import (
	"context"
	"encoding/json"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/delivery/worker"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/pkg/validator"
	"github.com/hibiken/asynq"
)

type LogLovePayload struct {
	// TODO: add fields for payload
}

type LogLove model.Job[LogLovePayload]

type logLove struct {
	client    *db.Client
	domain    *domain_entries.Domain
	validator validator.Validator
	worker    *worker.Client
}

func NewLogLove(
	client *db.Client,
	domain *domain_entries.Domain,
	validator validator.Validator,
	worker *worker.Client,
) LogLove {
	return &logLove{
		client:    client,
		domain:    domain,
		validator: validator,
		worker:    worker,
	}
}

func (j *logLove) Key() string {
	return "job:log_love"
}

func (j *logLove) Enqueue(ctx context.Context, payload LogLovePayload) (*asynq.TaskInfo, error) {
	if err := j.validator.Validate(payload); err != nil {
		return nil, err
	}

	data, _ := json.Marshal(payload)
	task := asynq.NewTask(j.Key(), data)
	return j.worker.EnqueueContext(ctx, task, asynq.Queue("critical"))
}

func (j *logLove) Handle(ctx context.Context, payload LogLovePayload) error {
	// TODO: implement job here

	return nil
}
