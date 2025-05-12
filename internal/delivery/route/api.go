package route

import (
	"github.com/alfariiizi/go-service/internal/delivery/route/handler"
	"github.com/alfariiizi/go-service/internal/service"
	"github.com/labstack/echo/v4"
)

type httpApi struct {
	httpRoute HttpRoute

	services service.Services
}

func NewHttpApi(
	echo *echo.Echo,
	services service.Services,
) HttpApi {
	return &httpApi{
		httpRoute: NewHttpRoute(echo),
		services:  services,
	}
}

func (h *httpApi) RegisterHandler() error {
	r := h.httpRoute

	r.RouteGroup("/api", func(r HttpRoute) {
		common := handler.NewCommonHandler()
		r.GET("/health", common.GetHealth)
		r.GET("/ping", common.GetPing)

		r.RouteGroup("/users", func(r HttpRoute) {
			userHandler := handler.NewUserHandler(h.services.User)
			r.GET("", userHandler.GetAllUsers)
			r.GET(":id", userHandler.GetUserByID)
			r.POST("", userHandler.CreateUser)
			r.PUT(":id", userHandler.UpdateUser)
			r.DELETE(":id", userHandler.DeleteUser)
		})
	})

	return nil
}
