package usecase

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/types"
)

type UserRegisterInput struct {
	// TODO: Define fields
}
type UserRegisterOutput struct {
	// TODO: Define fields
}
type UserRegister model.Usecase[UserRegisterInput, UserRegisterOutput]

type userRegister struct {
	client *repository.Client
}

func NewUserRegister(repo *repository.Client) UserRegister {
	return &userRegister{
		client: repo,
	}
}

func (uc *userRegister) Execute(ctx context.Context, input UserRegisterInput) types.Result[UserRegisterOutput]  {
	// TODO: Implement logic

	return types.Ok(UserRegisterOutput{})
}
