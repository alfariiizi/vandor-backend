package system_service

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/usecase"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
)

type HealthInput struct {
	// TODO: Define fields
}
type HealthOutput struct {
	Message string
}
type Health model.Service[HealthInput, HealthOutput]

type health struct {
	client  *db.Client
	usecase *usecase.Usecases
}

func NewHealth(
	repo *db.Client,
	usecase *usecase.Usecases,
) Health {
	return &health{
		client:  repo,
		usecase: usecase,
	}
}

func (s *health) Execute(ctx context.Context, input HealthInput) (*HealthOutput, error) {
	// TODO: Implement logic
	return &HealthOutput{
		Message: "Service is healthy",
	}, nil
}
