package http

import (
	"log"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/route"
	http "github.com/alfariiizi/vandor/internal/delivery/http/server"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(
		chi.NewMux,
		http.NewHttpServer,
		api.NewHttpApi,
	),
	route.Module,
	fx.Invoke(
		func(s *http.HttpServer) {
			cfg := config.GetConfig()
			log.Println("HTTP server is running on port ", cfg.HTTP.Port)
		},
	),
)
