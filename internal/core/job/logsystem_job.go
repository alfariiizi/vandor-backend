package job

import (
	"context"
	"encoding/json"
	"log"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/delivery/worker"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/pkg/validator"
	"github.com/hibiken/asynq"
)

type LogSystemPayload struct {
	// TODO: add fields
	Message string
}

type LogSystem model.Job[LogSystemPayload]

type logSystem struct {
	client    *db.Client
	domain    *domain_entries.Domain
	validator validator.Validator
	worker    *worker.Client
}

func NewLogSystem(
	client *db.Client,
	domain *domain_entries.Domain,
	validator validator.Validator,
	worker *worker.Client,
) LogSystem {
	return &logSystem{
		client:    client,
		domain:    domain,
		validator: validator,
		worker:    worker,
	}
}

func (j *logSystem) Key() string {
	return "job:log_system"
}

func (j *logSystem) Enqueue(ctx context.Context, payload LogSystemPayload) (*asynq.TaskInfo, error) {
	if err := j.validator.Validate(payload); err != nil {
		return nil, err
	}

	data, _ := json.Marshal(payload)
	task := asynq.NewTask(j.Key(), data)
	return j.worker.EnqueueContext(ctx, task, asynq.Queue("default"))
}

func (j *logSystem) Handle(ctx context.Context, payload LogSystemPayload) error {
	// TODO: implement
	log.Println("Handling log system job", "payload", payload)

	return nil
}
