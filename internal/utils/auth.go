package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type ExtractJWTOutput struct {
	UserID    uuid.UUID `json:"user_id"`
	SessionID string    `json:"session_id"`
	TenantID  uuid.UUID `json:"tenant_id"`
	Role      string    `json:"role"`
	Expiry    time.Time `json:"expiry"`
	IssuedAt  time.Time `json:"issued_at"`
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
	userUUID, err := IDParser(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %v", err)
	}

	sessionID, ok := getTokenIndex[string](token, "sid")
	if !ok {
		return nil, fmt.Errorf("failed to extract session ID from token")
	}

	tenantID, ok := getTokenIndex[string](token, "tid")
	if !ok {
		return nil, fmt.Errorf("failed to extract tenant ID from token")
	}
	tenantUUID, err := IDParser(tenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tenant ID: %v", err)
	}

	role, ok := getTokenIndex[string](token, "role")
	if !ok {
		return nil, fmt.Errorf("failed to extract role from token")
	}

	expiry, ok := getTokenIndex[time.Time](token, "exp")
	if !ok {
		return nil, fmt.Errorf("failed to extract expiry from token")
	}

	issuedAt, ok := getTokenIndex[time.Time](token, "iat")
	if !ok {
		return nil, fmt.Errorf("failed to extract issued at from token")
	}

	return &ExtractJWTOutput{
		UserID:    *userUUID,
		SessionID: sessionID,
		TenantID:  *tenantUUID,
		Role:      role,
		Expiry:    expiry,
		IssuedAt:  issuedAt,
	}, nil
}

func getTokenIndex[T any](token jwt.Token, key string) (T, bool) {
	var zero T
	value, found := token.Get(key)
	fmt.Println(key, ": ", value)
	if !found {
		return zero, false
	}
	result, ok := value.(T)
	if !ok {
		return zero, false
	}
	return result, true
}
