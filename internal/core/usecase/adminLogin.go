package usecase

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/types"
)

type AdminLoginInput struct {
	// TODO: Define fields
}
type AdminLoginOutput struct {
	// TODO: Define fields
}
type AdminLogin model.Usecase[AdminLoginInput, AdminLoginOutput]

type adminLogin struct {
	client *repository.Client
}

func NewAdminLogin(repo *repository.Client) AdminLogin {
	return &adminLogin{
		client: repo,
	}
}

func (uc *adminLogin) Execute(ctx context.Context, input AdminLoginInput) types.Result[AdminLoginOutput]  {
	// TODO: Implement logic

	return types.Ok(AdminLoginOutput{})
}
