package api

import (
	"fmt"
	"net/http"

	"github.com/alfariiizi/vandor/config"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type HttpApi struct {
	Api     huma.API
	BaseAPI huma.API
}

func NewHttpApi(
	router *chi.Mux,
) *HttpApi {
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(middleware.Logger)

	cfg := config.GetConfig()
	humaCfg := huma.DefaultConfig(cfg.App.Name, cfg.App.Version)
	humaCfg.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearerAuth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
			Description:  "Use the JWT token obtained from the login endpoint.",
		},
	}
	api := humachi.New(router, humaCfg)

	username := cfg.Docs.Username
	password := cfg.Docs.Password

	// Create a subrouter that requires authentication
	router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth(fmt.Sprintf("%s Docs", cfg.App.Name), map[string]string{
			username: password,
		}))

		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!doctype html>
			<html>
			  <head>
			    <title>API Reference</title>
			    <meta charset="utf-8" />
			    <meta name="viewport" content="width=device-width, initial-scale=1" />
			  </head>
			  <body>
			    <script
			      id="api-reference"
			      data-url="/openapi.json"
			      data-theme="light"
			      data-layout="sidebar"
			      data-auth="bearer"
			    ></script>
			    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
			  </body>
			</html>`))
		})
	})

	baseAPI := huma.NewGroup(api, "/api/v1")
	return &HttpApi{
		Api:     api,
		BaseAPI: baseAPI,
	}
}
