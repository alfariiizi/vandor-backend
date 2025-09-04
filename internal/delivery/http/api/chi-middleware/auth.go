package chimiddleware

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

var ignoredPaths []string = []string{"/api/admin/openapi.json", "/api/admin/docs"}

func ignoredAuthPath(r *http.Request) bool {
	return slices.Contains(ignoredPaths, r.URL.Path)
}

func AuthMiddleware(secretKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for openapi + docs
			if ignoredAuthPath(r) {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]

			// Parse & verify JWT
			token, err := jwt.Parse([]byte(tokenStr), jwt.WithKey(jwa.HS256, secretKey))
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			// Put token into context for later handlers
			ctx := context.WithValue(r.Context(), UserContextKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Extract user later
func GetJWTToken(r *http.Request) (jwt.Token, error) {
	token, ok := r.Context().Value(UserContextKey).(jwt.Token)
	if !ok {
		return nil, errors.New("no jwt token in context")
	}
	return token, nil
}
