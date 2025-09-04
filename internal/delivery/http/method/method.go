package method

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/delivery/http/api/middleware"
	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/danielgtaylor/huma/v2"
)

type Operation struct {
	Summary     string
	Description string
	Tags        []string
	BearerAuth  bool
	Tenant      bool
	Job         bool
	RoleAllowed []enum.UserRole
	Extensions  map[string]any
}

func generateBaseAPI[I, O any](api huma.API, path string, method string, operation Operation, handler func(context.Context, *I) (*O, error)) {
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

	if operation.Tenant {
		route.UseMiddleware(middleware.NewTenantMiddleware(route))
	}

	if operation.Job {
		route.UseMiddleware(middleware.NewJobMiddleware(route))
	}

	if len(operation.RoleAllowed) > 0 {
		route.UseMiddleware(middleware.UserRoleMiddleware(route, operation.RoleAllowed))
	}

	huma.Register(route, huma.Operation{
		OperationID: fmt.Sprintf("%s-%s", strings.ToLower(method), path),
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
	generateBaseAPI(api, path, http.MethodGet, operation, handler)
}

func POST[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodPost, operation, handler)
}

func PUT[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodPut, operation, handler)
}

func DELETE[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodDelete, operation, handler)
}

func PATCH[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodPatch, operation, handler)
}

func HEAD[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodHead, operation, handler)
}

func OPTIONS[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodOptions, operation, handler)
}

func TRACE[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodTrace, operation, handler)
}

func CONNECT[I any, O any](api huma.API, path string, operation Operation, handler func(context.Context, *I) (*O, error)) {
	generateBaseAPI(api, path, http.MethodConnect, operation, handler)
}
