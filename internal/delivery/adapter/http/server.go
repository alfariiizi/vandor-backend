package http

import (
	"context"
	"fmt"
	netHttp "net/http"

	"github.com/alfariiizi/go-service/config"
	serviceport "github.com/alfariiizi/go-service/internal/core/service/port"
	"github.com/alfariiizi/go-service/internal/delivery/adapter/http/handler"
	httpport "github.com/alfariiizi/go-service/internal/delivery/port/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type httpServer struct {
	echo *echo.Echo

	userService serviceport.UserService
}

// Setup sets up the HTTP server with the given configuration.
func NewHttpServer(
	lc fx.Lifecycle,
	userService serviceport.UserService,
) httpport.HttpServer {
	e := echo.New()
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
		echo:        e,
		userService: userService,
	}

	server.RegisterHandler()

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

func (h *httpServer) Start() error {
	cfg := config.GetConfig()
	h.echo.Logger.Fatal(h.echo.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
	return nil
}

func (h *httpServer) Stop(ctx context.Context) error {
	h.echo.Shutdown(ctx)
	return nil
}

func (h *httpServer) RegisterHandler() error {
	e := h.echo

	baseRoute := e.Group("/api")
	baseRoute.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	baseRoute.GET("/hello", func(c echo.Context) error {
		return handler.HelloWorldHandler(NewHttpContext(c))
	})
	baseRoute.GET("/hello/:id", func(c echo.Context) error {
		return handler.HelloWorldHandlerWithParams(NewHttpContext(c))
	})

	userRoute := baseRoute.Group("/user")
	userRoute.GET("", func(c echo.Context) error {
		return handler.GetAllUsersHandler(NewHttpContext(c), h.userService)
	})
	userRoute.GET(":id", func(c echo.Context) error {
		return handler.GetUserHandler(NewHttpContext(c), h.userService)
	})
	userRoute.POST("", func(c echo.Context) error {
		return handler.CreateUserHandler(NewHttpContext(c), h.userService)
	})
	userRoute.PUT(":id", func(c echo.Context) error {
		return handler.UpdateUserHandler(NewHttpContext(c), h.userService)
	})
	userRoute.DELETE(":id", func(c echo.Context) error {
		return handler.DeleteUserHandler(NewHttpContext(c), h.userService)
	})

	return nil
}
