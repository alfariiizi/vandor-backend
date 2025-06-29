package usecase

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/types"
)

type CreateSessionInput struct {
	// TODO: Define fields
}
type CreateSessionOutput struct {
	// TODO: Define fields
}
type CreateSession model.Usecase[CreateSessionInput, CreateSessionOutput]

type createSession struct {
	Client *repository.Client
}

func NewCreateSession(params createSession) CreateSession {
	return &params
}

func (uc *createSession) Execute(ctx context.Context, input CreateSessionInput) types.Result[CreateSessionOutput] {
	// TODO: Implement logic
	return types.Errf[CreateSessionOutput]("not implemented")
}
