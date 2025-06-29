package usecase

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/types"
)

type LoginInput struct {
	// TODO: Define fields
}
type LoginOutput struct {
	// TODO: Define fields
}
type Login model.Usecase[LoginInput, LoginOutput]

type login struct {
	client *repository.Client
}

func NewLogin(repo *repository.Client) Login {
	return &login{
		client: repo,
	}
}

func (uc *login) Execute(ctx context.Context, input LoginInput) types.Result[LoginOutput] {
	// TODO: Implement logic

	return types.Ok(LoginOutput{})
}
