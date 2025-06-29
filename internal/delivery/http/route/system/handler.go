package system_handler

import (
	"context"
	"log"

	"github.com/alfariiizi/go-service/internal/core/model"
	system_service "github.com/alfariiizi/go-service/internal/core/service/system"
	"github.com/alfariiizi/go-service/internal/delivery/http/api"
	"github.com/alfariiizi/go-service/internal/delivery/http/method"
	"github.com/alfariiizi/go-service/internal/utils"
	"go.uber.org/fx"
)

type params struct {
	fx.In

	*api.HttpApi
	HealthSvc system_service.Health
	PingSvc   system_service.Ping
}

func NewHandler(
	params params,
) model.Handler {
	handler := &params
	handler.RegisterRoutes()
	return handler
}

func (h *params) RegisterRoutes() {
	api := h.Api
	method.GET(api, "/health", method.Operation{
		Summary:     "Health",
		Description: "Check the health of the system",
		Tags:        []string{"System"},
		BearerAuth:  false,
	}, h.health)
	method.GET(api, "/ping", method.Operation{
		Summary:     "Ping",
		Description: "Ping the system to check if it's alive",
		Tags:        []string{"System"},
		BearerAuth:  false,
	}, h.ping)
}

func (h *params) health(ctx context.Context, input *struct{}) (*utils.OutputResponseData[system_service.HealthOutput], error) {
	resp := h.HealthSvc.Execute(ctx, system_service.HealthInput{})
	if resp.IsErr() {
		log.Println("Health check failed:", resp.Error())
		return nil, resp.Error()
	}
	return utils.GenerateOutputResponseData(resp.Unwrap()), nil
}

func (h *params) ping(ctx context.Context, input *struct{}) (*utils.OutputResponseData[system_service.PingOutput], error) {
	resp := h.PingSvc.Execute(ctx, system_service.PingInput{})
	if resp.IsErr() {
		return nil, resp.Error()
	}
	return utils.GenerateOutputResponseData(resp.Unwrap()), nil
}
