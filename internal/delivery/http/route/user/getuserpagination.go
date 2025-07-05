package user_handler

import (
	"context"

	"github.com/alfariiizi/go-service/internal/core/model"
	"github.com/alfariiizi/go-service/internal/core/service"
	user_service "github.com/alfariiizi/go-service/internal/core/service/user"
	"github.com/alfariiizi/go-service/internal/delivery/http/api"
	"github.com/alfariiizi/go-service/internal/delivery/http/method"
	"github.com/alfariiizi/go-service/internal/types"
	"github.com/danielgtaylor/huma/v2"
)

// NOTE:
// Hint Tags for input parameters
// @ref: https://huma.rocks/features/request-inputs
//
// Tag       | Description                           | Example
// -------------------------------------------------------------------
// path      | Name of the path parameter            | path:"thing-id"
// query     | Name of the query string parameter    | query:"q"
// header    | Name of the header parameter          | header:"Authorization"
// cookie    | Name of the cookie parameter          | cookie:"session"
// required  | Mark a query/header param as required | required:"true"

type GetUserPaginationInput struct {
	// Example GET input
	// ID    string `path:"id" doc:"ID of the item" example:"123"`
	// Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type GetUserPaginationData struct {
	// Example response data
	ID          string `json:"id" example:"123"`
	Name        string `json:"name" example:"Book"`
	Description string `json:"description" example:"A great book"`
}

type GetUserPaginationHandler model.HTTPHandler[GetUserPaginationInput, types.OutputResponsePagination[[]GetUserPaginationData]]

type getUserPagination struct {
	api     huma.API
	service *service.Services
}

func NewGetUserPagination(
	api *api.HttpApi,
	service *service.Services,
) GetUserPaginationHandler {
	h := &getUserPagination{
		api:     api.BaseAPI,
		service: service,
	}
	h.RegisterRoutes()
	return h
}

func (h *getUserPagination) RegisterRoutes() {
	api := h.api
	method.GET(api, "/users", method.Operation{
		Summary:     "GetUserPagination",
		Description: "GetUserPagination handler",
		Tags:        []string{"User"},
		BearerAuth:  false,
	}, h.Handler)
}

func (h *getUserPagination) Handler(ctx context.Context, input *GetUserPaginationInput) (*types.OutputResponsePagination[[]GetUserPaginationData], error) {
	res := h.service.User.GetUserPagination.Execute(ctx, user_service.GetUserPaginationInput{})
	if res.IsErr() {
		return nil, res.Error()
	}
	data := types.GenerateOutputPaginationData(
		[]GetUserPaginationData{
			{
				ID:          "1",
				Name:        "John Doe",
				Description: "A sample user",
			},
			{
				ID:          "2",
				Name:        "Jane Doe",
				Description: "Another sample user",
			},
		},
		1, 10, 5, 40,
	)
	return &data, nil
}
