package system_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/types"
	"go.uber.org/fx"
)

type PingInput struct {
	// TODO: Define fields
}
type PingOutput struct {
	Message string `json:"message"`
}
type Ping model.Service[PingInput, PingOutput]

type ping struct {
	fx.In
	Client *repository.Client
}

func NewPing(param ping) Ping {
	return &param
}

func (uc *ping) Execute(ctx context.Context, input PingInput) types.Result[PingOutput] {
	return types.Ok(PingOutput{
		Message: "pong",
	})
}
