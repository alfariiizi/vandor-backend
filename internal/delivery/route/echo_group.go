package route

import (
	httpctx "github.com/alfariiizi/go-service/internal/delivery/route/context"
	"github.com/labstack/echo/v4"
)

type httpGroup struct {
	group *echo.Group
}

func NewHttpRouteGroup(e *echo.Group) HttpRoute {
	return &httpGroup{
		group: e,
	}
}

func (h *httpGroup) RouteGroup(path string, fn func(r HttpRoute)) {
	group := h.group.Group(path)
	route := NewHttpRouteGroup(group)
	fn(route)
}

func (h *httpGroup) WithMiddleware(middleware ...func(next func(c httpctx.HttpContext) error) func(c httpctx.HttpContext) error) {
	// for _, m := range middleware {
	// 	h.group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 		return func(c echo.Context) error {
	// 			ctx := httpctx.NewHttpContext(c)
	// 			return m(func(c httpctx.HttpContext) error {
	// 				return next(c.(*httpctx.httpContext).context)
	// 			})(ctx)
	// 		}
	// 	})
	// }
}

func (h *httpGroup) GET(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.GET(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) POST(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.POST(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) PUT(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.PUT(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) DELETE(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.DELETE(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) PATCH(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.PATCH(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) OPTIONS(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.OPTIONS(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) HEAD(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.HEAD(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) CONNECT(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.CONNECT(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}

func (h *httpGroup) TRACE(path string, handlerFunc func(c httpctx.HttpContext) error) {
	h.group.TRACE(path, func(c echo.Context) error {
		ctx := httpctx.NewHttpContext(c)
		return handlerFunc(ctx)
	})
}
