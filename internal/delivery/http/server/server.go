package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alfariiizi/vandor/config"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"go.uber.org/fx"

	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type HttpServer struct {
	Router *chi.Mux
}

func NewHttpServer(
	lc fx.Lifecycle,
	router *chi.Mux,
	api *api.HttpApi,
) *HttpServer {
	// router := chi.NewMux()
	cfg := config.GetConfig()

	// api.RegisterHandler()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Http.Port),
		Handler: router,
	}

	// Serve static files from ./public at root path
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./storage/public"))))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					log.Fatal("Failed to start server: ", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Shutdown(ctx)
			return nil
		},
	})

	return &HttpServer{
		Router: router,
	}
}
