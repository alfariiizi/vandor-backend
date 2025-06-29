package method

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/go-service/config"
	"github.com/alfariiizi/go-service/internal/delivery/http/api/middleware"
)

type Operation struct {
	Summary     string
	Description string
	Tags        []string
	BearerAuth  bool
	Extensions  map[string]any
}

func generateBaseApi[I, O any](api huma.API, path string, method string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	route := huma.NewGroup(api, "")
	var security []map[string][]string
	cfg := config.GetConfig()
	if operation.BearerAuth {
		security = []map[string][]string{
			{
				"bearerAuth": {},
			},
		}
		route.UseMiddleware(middleware.NewAuthMiddleware(route, cfg.Auth.SecretKey))
	} else {
		security = []map[string][]string{
			{},
		}
	}
	huma.Register(route, huma.Operation{
		OperationID: fmt.Sprintf("%s-%s", method, path),
		Method:      method,
		Path:        path,
		Summary:     operation.Summary,
		Description: operation.Description,
		Tags:        operation.Tags,
		Security:    security,
		Extensions:  operation.Extensions,
	}, handler)
}

func GET[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodGet, operation, handler)
}

func POST[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodPost, operation, handler)
}

func PUT[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodPut, operation, handler)
}

func DELETE[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodDelete, operation, handler)
}

func PATCH[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodPatch, operation, handler)
}

func HEAD[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodHead, operation, handler)
}

func OPTIONS[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodOptions, operation, handler)
}

func TRACE[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodTrace, operation, handler)
}

func CONNECT[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseApi(api, path, http.MethodConnect, operation, handler)
}
