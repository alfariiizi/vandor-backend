package system_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/usecase"
	"github.com/alfariiizi/go-service/internal/types"
)

type PingInput struct {
	// TODO: Define fields
}
type PingOutput struct {
	Message string
}
type Ping model.Service[PingInput, PingOutput]

type ping struct {
	client  *repository.Client
	usecase *usecase.Usecases
}

func NewPing(
	repo *repository.Client,
	usecase *usecase.Usecases,
) Ping {
	return &ping{
		client:  repo,
		usecase: usecase,
	}
}

func (s *ping) Execute(ctx context.Context, input PingInput) types.Result[PingOutput] {
	return types.Ok(PingOutput{
		Message: "Pong",
	})
}
