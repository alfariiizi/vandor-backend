package system_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/types"
	"go.uber.org/fx"
)

type HealthInput struct {
	// TODO: Define fields
}
type HealthOutput struct {
	Status string `json:"status"`
}
type Health model.Service[HealthInput, HealthOutput]

type health struct {
	fx.In
	Client *repository.Client
}

func NewHealth(params health) Health {
	return &params
}

func (uc *health) Execute(ctx context.Context, input HealthInput) types.Result[HealthOutput] {
	return types.Ok(HealthOutput{
		Status: "healthy",
	})
	// return &HealthOutput{
	// 	Status: "healthy",
	// }, nil
}
