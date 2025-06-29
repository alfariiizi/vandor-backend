package user_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/types"
)

type GetAllUserAbcInput struct {
	// TODO: Define fields
}
type GetAllUserAbcOutput struct {
	// TODO: Define fields
}
type GetAllUserAbc model.Service[GetAllUserAbcInput, GetAllUserAbcOutput]

type getAllUserAbc struct {
	client *repository.Client
}

func NewGetAllUserAbc(repo *repository.Client) GetAllUserAbc {
	return &getAllUserAbc{
		client: repo,
	}
}

func (uc *getAllUserAbc) Execute(ctx context.Context, input GetAllUserAbcInput) types.Result[GetAllUserAbcOutput]  {
	// TODO: Implement logic

	return types.Ok(GetAllUserAbcOutput{})
}
