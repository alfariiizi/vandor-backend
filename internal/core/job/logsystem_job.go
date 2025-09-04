package job

import (
	"context"
	"encoding/json"
	"strings"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/delivery/worker"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/pkg/logger"
	"github.com/alfariiizi/vandor/internal/pkg/validator"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/danielgtaylor/huma/v2"
	"github.com/hibiken/asynq"
)

type LogSystemPayload struct {
	// TODO: add fields
	Message string `json:"message" required:"true" validate:"required"`
}

type LogSystem model.Job[LogSystemPayload]

type LogSystemHTTPInput struct {
	JobSecret string           `header:"X-Job-Secret" required:"true"`
	Body      LogSystemPayload `json:"body" contentType:"application/json"`
}

type logSystem struct {
	api       huma.API
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
	inspector *worker.Inspector,
	api *api.HttpApi,
) LogSystem {
	return &logSystem{
		client:    client,
		domain:    domain,
		validator: validator,
		worker:    worker,
		api:       api.JobAPI,
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
	log := logger.Get()
	log.Info().
		Str("message", payload.Message).
		Msg("Log System Job Executed")

	return nil
}

func (j *logSystem) HTTPRegisterRoute() {
	path := "/" + strings.Split(j.Key(), ":")[1]

	method.POST(j.api, path, method.Operation{
		Summary:     "Log System Job",
		Description: "Enqueue a job to log system messages",
		Tags:        []string{"Job"},
		Job:         true,
	}, func(ctx context.Context, input *LogSystemHTTPInput) (*model.JobHTTPHandlerResponse, error) {
		taskInfo, err := j.Enqueue(ctx, input.Body)
		if err != nil {
			return nil, err
		}
		return (*model.JobHTTPHandlerResponse)(types.GenerateOutputResponseData(model.JobHTTPHandlerData{
			TaskID: taskInfo.ID,
		})), nil
	})
}
