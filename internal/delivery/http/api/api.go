package api

import (
	"fmt"
	"net/http"

	"github.com/alfariiizi/vandor/internal/config"
	chimiddleware "github.com/alfariiizi/vandor/internal/delivery/http/api/chi-middleware"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/rest"
	"github.com/alfariiizi/vandor/internal/infrastructure/sse"
	"github.com/alfariiizi/vandor/internal/monitoring"
	"github.com/alfariiizi/vandor/internal/pkg/logger"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HttpApi struct {
	Api     huma.API
	BaseAPI huma.API
	JobAPI  huma.API
}

func NewHttpApi(
	router *chi.Mux,
	register *prometheus.Registry,
	sseMgr *sse.Manager,
	dbClient *db.Client,
) *HttpApi {
	cfg := config.GetConfig()
	logger := logger.Get()

	// --- existing middleware & routes (mostly unchanged) ---
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Tenant-ID"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	chimiddleware.RegisterZerolog(router, *logger, "/api", []string{"/api/admin/docs", "/api/admin/openapi.json"})

	router.Group(func(r chi.Router) {
		monit := Monitoring(register)
		r.Use(monitoring.MetricsMiddleware(monit))
		httpRequests := prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"path"},
		)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			httpRequests.WithLabelValues("/").Inc()
			w.Write([]byte("Hello Vandor 🚀"))
		})
		r.Handle("/metrics", promhttp.HandlerFor(register, promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		}))
	})

	router.Get("/events", sseMgr.ServeHTTP)

	router.Group(func(r chi.Router) {
		r.Use(chimiddleware.AuthMiddleware([]byte(cfg.Auth.SecretKey)))

		srv, err := rest.NewServer(dbClient, &rest.ServerConfig{
			BasePath: "/api/admin",
		})
		if err != nil {
			panic(fmt.Sprintf("failed to create entrest server: %v", err))
		}

		r.Mount("/", srv.Handler())
	})

	router.Group(func(r chi.Router) {
		username := cfg.Docs.Username
		password := cfg.Docs.Password
		r.Use(middleware.BasicAuth(fmt.Sprintf("%s ERD", cfg.App.Name), map[string]string{
			username: password,
		}))
		r.Mount("/erd", db.ServeEntviz())
	})

	// Huma API + docs config (unchanged)
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

	// Basic-auth-protected docs UI (unchanged)
	router.Group(func(r chi.Router) {
		username := cfg.Docs.Username
		password := cfg.Docs.Password
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

	// --- Huma groups as before ---
	baseAPI := huma.NewGroup(api, "/api")
	jobAPI := huma.NewGroup(baseAPI, "/jobs")

	return &HttpApi{
		Api:     api,
		BaseAPI: baseAPI,
		JobAPI:  jobAPI,
	}
}

func Monitoring(register *prometheus.Registry) *monitoring.Metrics {
	m := monitoring.NewMetrics(register)
	return m
}
