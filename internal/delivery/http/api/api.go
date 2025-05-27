package api

import (
	"github.com/alfariiizi/go-service/internal/delivery/http/route/handler"
	"github.com/alfariiizi/go-service/internal/service"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

type HttpApi struct {
	services service.Services
	Api      huma.API
}

func NewHttpApi(
	router *chi.Mux,
	services service.Services,
) *HttpApi {
	api := humachi.New(router, huma.DefaultConfig("Go Services", "1.0.0"))

	return &HttpApi{
		Api:      api,
		services: services,
	}
}

func (h *HttpApi) RegisterHandler() error {
	base := huma.NewGroup(h.Api, "/api")

	commonHandler := handler.NewCommonHandler()
	GET(base, "/health", Operation{
		Summary:     "Get Health",
		Description: "Get the health status of the service",
		Tags:        []string{"Common"},
	}, commonHandler.GetHealth)
	GET(base, "/ping", Operation{
		Summary:     "Ping the service",
		Description: "Ping the service to check if it is alive",
		Tags:        []string{"Common"},
	}, commonHandler.GetPing)

	return nil
}
