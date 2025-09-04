package middleware

import (
	"net/http"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/danielgtaylor/huma/v2"
)

// WithTenantMiddleware extracts the X-Tenant-ID from headers and stores it in context
func NewJobMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		secret := config.GetConfig().Jobs.HTTPHeaderSecret

		jobSecret := ctx.Header("X-Job-Secret")
		if jobSecret == "" {
			huma.WriteErr(api, ctx, http.StatusBadRequest, "Missing X-Job-Secret header")
			return
		}

		if jobSecret != secret {
			huma.WriteErr(api, ctx, http.StatusForbidden, "Invalid X-Job-Secret header")
			return
		}

		next(ctx)
	}
}
