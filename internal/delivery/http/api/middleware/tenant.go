package middleware

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// WithTenantMiddleware extracts the X-Tenant-ID from headers and stores it in context
func NewTenantMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		tenantID := ctx.Header("X-Tenant-ID")
		if tenantID == "" {
			huma.WriteErr(api, ctx, http.StatusBadRequest, "Missing X-Tenant-ID header")
			return
		}

		ctx = huma.WithValue(ctx, "tenant", tenantID)

		// Optional: verify tenant access
		// token, ok := ctx.Value("user").(jwt.Token)
		// token, ok := ctx.Context().Value("user").(jwt.Token)
		// if !ok {
		// 	huma.WriteErr(api, ctx, http.StatusUnauthorized, "User context missing or invalid")
		// 	return
		// }

		// Get user tenant list from token (example claim: "tenants": ["tenant_1", "tenant_2"])
		// tenantsClaim, ok := token.Get("tenants")
		// if !ok {
		// 	huma.WriteErr(api, ctx, http.StatusForbidden, "No tenant access defined in token")
		// 	return
		// }

		// tenants, ok := tenantsClaim.([]any) // JWT library uses []any for JSON arrays
		// if !ok {
		// 	huma.WriteErr(api, ctx, http.StatusForbidden, "Invalid tenant claim format")
		// 	return
		// }
		//
		// hasAccess := false
		// for _, t := range tenants {
		// 	if t.(string) == tenantID {
		// 		hasAccess = true
		// 		break
		// 	}
		// }
		//
		// if !hasAccess {
		// 	huma.WriteErr(api, ctx, http.StatusForbidden, "User does not have access to tenant")
		// 	return
		// }
		//
		// // Add tenant ID to context
		// ctx = huma.WithValue(ctx, "tenantID", tenantID)
		next(ctx)
	}
}
