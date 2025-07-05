package user_service

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/repository"
	"github.com/alfariiizi/go-service/internal/core/usecase"
	"github.com/alfariiizi/go-service/internal/types"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserPaginationInput struct {
	// TODO: Define fields
}
type GetUserPaginationOutput model.PaginationOutput[[]User]

type GetUserPagination model.Service[GetUserPaginationInput, GetUserPaginationOutput]

type getUserPagination struct {
	client  *repository.Client
	usecase *usecase.Usecases
}

func NewGetUserPagination(
	repo *repository.Client,
	usecase *usecase.Usecases,
) GetUserPagination {
	return &getUserPagination{
		client:  repo,
		usecase: usecase,
	}
}

func (s *getUserPagination) Execute(ctx context.Context, input GetUserPaginationInput) types.Result[GetUserPaginationOutput] {
	// TODO: Implement logic

	return types.Ok(GetUserPaginationOutput{
		Data: []User{
			{
				ID:    "1",
				Name:  "John Doe",
				Email: "johndoe@mail.com",
			},
			{
				ID:    "2",
				Name:  "Jane Doe",
				Email: "janedoe@mail.com",
			},
		},
		Meta: model.PaginationMeta{
			Page:       1,
			PerPage:    10,
			TotalPages: 5,
			TotalItems: 42,
		},
	})
}
