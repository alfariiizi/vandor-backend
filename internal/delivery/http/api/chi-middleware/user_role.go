package chimiddleware

import (
	"net/http"

	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/alfariiizi/vandor/internal/utils"
)

var allowedRoles = []enum.UserRole{
	enum.UserRoleAdmin,
	enum.UserRoleSuperAdmin,
}

// UserRoleMiddlewareChi creates a Chi middleware that validates user role
func UserRoleMiddlewareChi() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ignoredAuthPath(r) {
				next.ServeHTTP(w, r)
				return
			}

			token, err := utils.ExtractJWT(r.Context())
			if err != nil {
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
				return
			}

			userRole := token.Role
			if userRole == "" {
				http.Error(w, "User role not found in token", http.StatusForbidden)
				return
			}

			hasRole := false
			for _, role := range allowedRoles {
				if userRole == role.Label() {
					hasRole = true
					break
				}
			}

			if !hasRole {
				http.Error(w, "User does not have required role", http.StatusForbidden)
				return
			}

			// Pass through if authorized
			next.ServeHTTP(w, r)
		})
	}
}
