package route

import (
	httpctx "github.com/alfariiizi/go-service/internal/delivery/route/context"
	"github.com/labstack/echo/v4"
)

type httpRoute struct {
	echo *echo.Echo
}

func NewHttpRoute(e *echo.Echo) HttpRoute {
	return &httpRoute{
		echo: e,
	}
}

func (h *httpRoute) RouteGroup(path string, fn func(r HttpRoute)) {
	group := h.echo.Group(path)
	fn(NewHttpRouteGroup(group))
}

func (h *httpRoute) WithMiddleware(middleware ...func(next func(c httpctx.HttpContext) error) func(c httpctx.HttpContext) error) {
	// for _, m := range middleware {
	// 	h.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 		return func(c echo.Context) error {
	// 			ctx := httpctx.NewHttpContext(c)
	// 			return m(func(c httpctx.HttpContext) error {
	// 				return next(c.(*httpContext).context)
	// 			})(ctx)
	// 		}
	// 	})
	// }
}

func (h *httpRoute) GET(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.GET(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) POST(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.POST(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) PUT(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.PUT(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) DELETE(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.DELETE(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) PATCH(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.PATCH(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) OPTIONS(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.OPTIONS(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) HEAD(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.HEAD(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) CONNECT(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.CONNECT(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpRoute) TRACE(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.echo.TRACE(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}
