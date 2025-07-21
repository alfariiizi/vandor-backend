package http

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/alfariiizi/vandor/config"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/route"
	http "github.com/alfariiizi/vandor/internal/delivery/http/server"
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
			log.Println("HTTP server is running on port ", cfg.Http.Port)
		},
	),
)
