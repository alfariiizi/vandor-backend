package utils

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwt"
)

type ExtractJWTOutput struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	Expiry    int64  `json:"expiry"`
	IssuedAt  int64  `json:"issued_at"`
}

func ExtractJWT(ctx context.Context) (*ExtractJWTOutput, error) {
	token, found := ctx.Value("user").(jwt.Token)
	if !found {
		return nil, fmt.Errorf("token is not valid")
	}
	fmt.Println("Token", token)

	userID, ok := getTokenIndex[string](token, "sub")
	if !ok {
		return nil, fmt.Errorf("failed to extract user ID from token")
	}

	sessionID, ok := getTokenIndex[string](token, "sid")
	if !ok {
		return nil, fmt.Errorf("failed to extract session ID from token")
	}

	expiry, ok := getTokenIndex[int64](token, "exp")
	if !ok {
		return nil, fmt.Errorf("failed to extract expiry from token")
	}

	issuedAt, ok := getTokenIndex[int64](token, "iat")
	if !ok {
		return nil, fmt.Errorf("failed to extract issued at from token")
	}

	return &ExtractJWTOutput{
		UserID:    userID,
		SessionID: sessionID,
		Expiry:    expiry,
		IssuedAt:  issuedAt,
	}, nil
}

func getTokenIndex[T any](token jwt.Token, key string) (T, bool) {
	var zero T
	value, found := token.Get(key)
	if !found {
		return zero, false
	}
	result, ok := value.(T)
	if !ok {
		return zero, false
	}
	return result, true
}
