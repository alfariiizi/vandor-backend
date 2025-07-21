package auth_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	auth_service "github.com/alfariiizi/vandor/internal/core/service/auth"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/danielgtaylor/huma/v2"
)

type RefreshPayload struct {
	RefreshToken string `json:"refresh_token" required:"true" example:"your-refresh-token" doc:"The refresh token to use for refreshing the session"`
}

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

type RefreshInput struct {
	// JSON body for POST
	Body RefreshPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type RefreshOutput types.OutputResponseData[RefreshData]

type RefreshData *auth_service.RefreshOutput

type RefreshHandler model.HTTPHandler[RefreshInput, RefreshOutput]

type refresh struct {
	api     huma.API
	service *service.Services
}

func NewRefresh(
	api *api.HttpApi,
	service *service.Services,
) RefreshHandler {
	h := &refresh{
		api:     api.BaseAPI,
		service: service,
	}
	h.RegisterRoutes()
	return h
}

func (h *refresh) RegisterRoutes() {
	api := h.api
	method.POST(api, "/refresh", method.Operation{
		Summary:     "Refresh",
		Description: "Refresh handler",
		Tags:        []string{"Auth"},
		BearerAuth:  false,
	}, h.Handler)
}

func (h *refresh) GenerateResponse(data RefreshData) *RefreshOutput {
	return (*RefreshOutput)(types.GenerateOutputResponseData(data))
}

func (h *refresh) Handler(ctx context.Context, input *RefreshInput) (*RefreshOutput, error) {
	payload := input.Body
	res, err := h.service.Auth.Refresh.Execute(ctx, auth_service.RefreshInput{
		RefreshToken: payload.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(RefreshData(res)), nil
}
