package system_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/usecase"
	"github.com/alfariiizi/go-service/internal/types"
)

type HealthInput struct {
	// TODO: Define fields
}
type HealthOutput struct {
	Message string
}
type Health model.Service[HealthInput, HealthOutput]

type health struct {
	client  *repository.Client
	usecase *usecase.Usecases
}

func NewHealth(
	repo *repository.Client,
	usecase *usecase.Usecases,
) Health {
	return &health{
		client:  repo,
		usecase: usecase,
	}
}

func (s *health) Execute(ctx context.Context, input HealthInput) types.Result[HealthOutput] {
	// TODO: Implement logic

	return types.Ok(HealthOutput{
		Message: "Service is healthy",
	})
}
