package middleware

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func NewAuthMiddleware(api huma.API, secretKey string) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		tokenStr := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
		if tokenStr == "" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		// Parse and validate token using shared secret
		token, err := jwt.ParseString(tokenStr,
			jwt.WithKey(jwa.HS256, []byte(secretKey)),
			jwt.WithValidate(true),
		)
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Save token to context for handler use
		ctx = huma.WithValue(ctx, "user", token)
		next(ctx)
	}
}

// import (
// 	"context"
// 	"net/http"
// 	"strings"
// 	"time"
//
// 	"github.com/danielgtaylor/huma/v2"
// 	"github.com/lestrrat-go/jwx/v2/jwk"
// 	"github.com/lestrrat-go/jwx/v2/jwt"
// )
//
// func NewAuthMiddleware(api huma.API, jwksURL, expectedIssuer, expectedAudience string) func(ctx huma.Context, next func(huma.Context)) {
// 	keySet := NewJWKSet(jwksURL)
//
// 	return func(ctx huma.Context, next func(huma.Context)) {
// 		tokenStr := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
// 		if tokenStr == "" {
// 			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing Authorization header")
// 			return
// 		}
//
// 		token, err := jwt.ParseString(tokenStr,
// 			jwt.WithKeySet(keySet),
// 			jwt.WithValidate(true),
// 			jwt.WithIssuer(expectedIssuer),
// 			jwt.WithAudience(expectedAudience),
// 		)
// 		if err != nil {
// 			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid or expired token")
// 			return
// 		}
//
// 		// Inject token into context
// 		ctx = huma.WithValue(ctx, "user", token)
// 		next(ctx)
// 	}
// }
//
// func NewJWKSet(jwkURL string) jwk.Set {
// 	jwkCache := jwk.NewCache(context.Background())
//
// 	err := jwkCache.Register(jwkURL, jwk.WithMinRefreshInterval(10*time.Minute))
// 	if err != nil {
// 		panic("failed to register JWK location: " + err.Error())
// 	}
//
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
//
// 	_, err = jwkCache.Refresh(ctx, jwkURL)
// 	if err != nil {
// 		panic("failed to fetch JWKs on startup: " + err.Error())
// 	}
//
// 	return jwk.NewCachedSet(jwkCache, jwkURL)
// }

//
// func NewAuthMiddleware(api huma.API, jwksURL, expectedIssuer, expectedAudience string) func(ctx huma.Context, next func(huma.Context)) {
// 	keySet := NewJWKSet(jwksURL)
//
// 	return func(ctx huma.Context, next func(huma.Context)) {
// 		var anyOfNeededScopes []string
// 		isAuthorizationRequired := false
//
// 		for _, opScheme := range ctx.Operation().Security {
// 			if scopes, ok := opScheme["myAuth"]; ok {
// 				anyOfNeededScopes = scopes
// 				isAuthorizationRequired = true
// 				break
// 			}
// 		}
//
// 		if !isAuthorizationRequired {
// 			next(ctx)
// 			return
// 		}
//
// 		token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
// 		if token == "" {
// 			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing Authorization header")
// 			return
// 		}
//
// 		parsed, err := jwt.ParseString(token,
// 			jwt.WithKeySet(keySet),
// 			jwt.WithValidate(true),
// 			jwt.WithIssuer(expectedIssuer),
// 			jwt.WithAudience(expectedAudience),
// 		)
// 		if err != nil {
// 			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
// 			return
// 		}
//
// 		// Optional: Scope enforcement
// 		if scopes, ok := parsed.Get("scopes"); ok {
// 			if scopeList, ok := scopes.([]any); ok {
// 				for _, scope := range scopeList {
// 					if scopeStr, ok := scope.(string); ok && slices.Contains(anyOfNeededScopes, scopeStr) {
// 						ctx = huma.WithValue(ctx, "user", parsed)
// 						next(ctx)
// 						return
// 					}
// 				}
// 			}
// 		}
//
// 		huma.WriteErr(api, ctx, http.StatusForbidden, "Insufficient scope")
// 	}
// }
