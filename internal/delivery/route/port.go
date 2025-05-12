package route

import httpctx "github.com/alfariiizi/go-service/internal/delivery/route/context"

type HttpRoute interface {
	// Group(fn func(r HttpRoute))
	RouteGroup(path string, fn func(r HttpRoute))
	WithMiddleware(middleware ...func(next func(c httpctx.HttpContext) error) func(c httpctx.HttpContext) error)

	GET(path string, handlerFunc func(c httpctx.HttpContext) error)
	POST(path string, handlerFunc func(c httpctx.HttpContext) error)
	PUT(path string, handlerFunc func(c httpctx.HttpContext) error)
	DELETE(path string, handlerFunc func(c httpctx.HttpContext) error)
	OPTIONS(path string, handlerFunc func(c httpctx.HttpContext) error)
	PATCH(path string, handlerFunc func(c httpctx.HttpContext) error)
	HEAD(path string, handlerFunc func(c httpctx.HttpContext) error)
	CONNECT(path string, handlerFunc func(c httpctx.HttpContext) error)
	TRACE(path string, handlerFunc func(c httpctx.HttpContext) error)
}

type HttpApi interface {
	RegisterHandler() error
}
