package system_service

import (
	"context"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/usecase"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
)

type PingInput struct {
	// TODO: Define fields
}
type PingOutput struct {
	Message string
}
type Ping model.Service[PingInput, PingOutput]

type ping struct {
	client  *db.Client
	usecase *usecase.Usecases
	domain  *domain_entries.Domain
}

func NewPing(
	repo *db.Client,
	usecase *usecase.Usecases,
	domain *domain_entries.Domain,
) Ping {
	return &ping{
		client:  repo,
		usecase: usecase,
		domain:  domain,
	}
}

func (s *ping) Execute(ctx context.Context, input PingInput) (*PingOutput, error) {
	// Return a successful ping response
	return &PingOutput{
		Message: "pong",
	}, nil
}
