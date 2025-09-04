package auth_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
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

type meInput struct{}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type meOutput types.OutputResponseData[meData]

type meData struct {
	*db.User
}

type meHandler model.HTTPHandler[meInput, meOutput]

type me struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func Newme(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) meHandler {
	h := &me{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *me) RegisterRoutes() {
	api := h.api
	method.GET(api, "/me", method.Operation{
		Summary:     "Get Authenticated User",
		Description: "Retrieve information about the authenticated user.",
		Tags:        []string{"Auth"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *me) GenerateResponse(data meData) *meOutput {
	return (*meOutput)(types.GenerateOutputResponseData(data))
}

func (h *me) Handler(ctx context.Context, input *meInput) (*meOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}
	userOne, err := h.client.User.Get(ctx, token.UserID)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(meData{userOne}), nil
}
