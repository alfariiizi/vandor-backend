package system_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	system_service "github.com/alfariiizi/vandor/internal/core/service/system"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
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

type HealthInput struct{}

// NOTE:
// You can use this for output
// types.OutputResponseData[T]
// types.OutputResponseMessage

type HealthOutput types.OutputResponseMessage

type HealthHandler model.HTTPHandler[HealthInput, HealthOutput]

type health struct {
	api     huma.API
	service *service.Services
}

func NewHealth(
	api *api.HttpApi,
	service *service.Services,
) HealthHandler {
	h := &health{
		api:     api.BaseAPI,
		service: service,
	}
	h.RegisterRoutes()
	return h
}

func (h *health) RegisterRoutes() {
	api := h.api
	method.GET(api, "/health", method.Operation{
		Summary:     "Health",
		Description: "Health handler",
		Tags:        []string{"System"},
		BearerAuth:  false,
	}, h.Handler)
}

func (h *health) Handler(ctx context.Context, input *HealthInput) (*HealthOutput, error) {
	res, err := h.service.System.Health.Execute(ctx, system_service.HealthInput{})
	if err != nil {
		return nil, err
	}
	return (*HealthOutput)(types.GenerateOutputResponseMessage(res.Message)), nil
}
