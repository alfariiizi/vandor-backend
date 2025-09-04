package middleware

import (
	"net/http"

	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

func UserRoleMiddleware(api huma.API, roles []enum.UserRole) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		token, err := utils.ExtractJWT(ctx.Context())
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		// Check if user has any of the required roles
		userRole := token.Role
		if userRole == "" {
			huma.WriteErr(api, ctx, http.StatusForbidden, "User role not found in token")
			return
		}
		hasRole := false
		for _, role := range roles {
			if userRole == role.Label() {

				hasRole = true
				break
			}
		}

		if !hasRole {
			huma.WriteErr(api, ctx, http.StatusForbidden, "User does not have required role")
			return
		}

		next(ctx)
	}
}
