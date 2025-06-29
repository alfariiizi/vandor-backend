package auth_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/types"
)

type UserLoginInput struct {
	// TODO: Define fields
}
type UserLoginOutput struct {
	// TODO: Define fields
}
type UserLogin model.Service[UserLoginInput, UserLoginOutput]

type userLogin struct {
	client *repository.Client
}

func NewUserLogin(repo *repository.Client) UserLogin {
	return &userLogin{
		client: repo,
	}
}

func (uc *userLogin) Execute(ctx context.Context, input UserLoginInput) types.Result[UserLoginOutput]  {
	// TODO: Implement logic

	return types.Ok(UserLoginOutput{})
}
