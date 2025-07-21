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

type PingInput struct{}

// NOTE:
// You can use this for output:
// types.OutputResponseData[T]
// types.OutputResponseMessage

type PingOutput types.OutputResponseData[PingData]

type PingData struct {
	Message string `json:"message" doc:"Ping response message" example:"pong"`
}

type PingHandler model.HTTPHandler[PingInput, PingOutput]

type ping struct {
	api     huma.API
	service *service.Services
}

func NewPing(
	api *api.HttpApi,
	service *service.Services,
) PingHandler {
	h := &ping{
		api:     api.BaseAPI,
		service: service,
	}
	h.RegisterRoutes()
	return h
}

func (h *ping) RegisterRoutes() {
	api := h.api
	method.GET(api, "/ping", method.Operation{
		Summary:     "Ping",
		Description: "Ping handler",
		Tags:        []string{"System"},
		BearerAuth:  false,
	}, h.Handler)
}

func (h *ping) Handler(ctx context.Context, input *PingInput) (*PingOutput, error) {
	res, err := h.service.System.Ping.Execute(ctx, system_service.PingInput{})
	if err != nil {
		return nil, err
	}
	return (*PingOutput)(types.GenerateOutputResponseData(PingData{
		Message: res.Message,
	})), nil
}
