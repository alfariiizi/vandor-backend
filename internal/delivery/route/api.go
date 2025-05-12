package route

import (
	"github.com/alfariiizi/go-service/internal/delivery/route/handler"
	"github.com/labstack/echo/v4"
)

type httpApi struct {
	httpRoute HttpRoute
}

func NewHttpApi(echo *echo.Echo) HttpApi {
	return &httpApi{
		httpRoute: NewHttpRoute(echo),
	}
}

func (h *httpApi) RegisterHandler() error {
	r := h.httpRoute

	r.RouteGroup("/api", func(r HttpRoute) {
		common := handler.NewCommonHandler()
		r.GET("/health", common.GetHealth)
		r.GET("/ping", common.GetPing)
	})

	return nil
}
