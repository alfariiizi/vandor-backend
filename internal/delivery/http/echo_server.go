package http

import (
	"context"
	"fmt"
	netHttp "net/http"

	"github.com/alfariiizi/go-service/config"
	"github.com/alfariiizi/go-service/internal/delivery/route"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type httpServer struct {
	echo *echo.Echo
}

// Setup sets up the HTTP server with the given configuration.
func NewHttpServer(
	lc fx.Lifecycle,
	e *echo.Echo,
	api route.HttpApi,
) HttpServer {
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	server := &httpServer{
		echo: e,
	}

	api.RegisterHandler()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.Start(); err != nil {
					e.Logger.Fatal("Failed to start server: ", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Stop(ctx); err != nil {
				e.Logger.Fatal("Failed to stop server: ", err)
			}
			return nil
		},
	})

	return server
}

func (h *httpServer) Start() error {
	cfg := config.GetConfig()
	h.echo.Logger.Fatal(h.echo.Start(fmt.Sprintf(":%d", cfg.Http.Port)))
	return nil
}

func (h *httpServer) Stop(ctx context.Context) error {
	h.echo.Shutdown(ctx)
	return nil
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(netHttp.StatusBadRequest, err.Error())
	}
	return nil
}
