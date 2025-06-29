package usecase

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/types"
)

type AdminRegisterInput struct {
	// TODO: Define fields
}
type AdminRegisterOutput struct {
	// TODO: Define fields
}
type AdminRegister model.Usecase[AdminRegisterInput, AdminRegisterOutput]

type adminRegister struct {
	client *repository.Client
}

func NewAdminRegister(repo *repository.Client) AdminRegister {
	return &adminRegister{
		client: repo,
	}
}

func (uc *adminRegister) Execute(ctx context.Context, input AdminRegisterInput) types.Result[AdminRegisterOutput]  {
	// TODO: Implement logic

	return types.Ok(AdminRegisterOutput{})
}
