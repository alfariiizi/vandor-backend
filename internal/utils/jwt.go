package utils

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type AccessToken struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func GenerateAccessToken(userID string, sessionID string, name string, email string, role string) (AccessToken, error) {
	cfg := config.GetConfig()
	duration := time.Now().Add(time.Duration(cfg.Auth.TokenDurationInMinutes) * time.Minute).Unix()
	token := jwt.New()
	token.Set("sub", userID)
	token.Set("sid", sessionID)
	token.Set("exp", duration)
	token.Set("name", name)
	token.Set("email", email)
	token.Set("role", role)
	token.Set("iat", time.Now().Unix())

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(cfg.Auth.SecretKey)))
	if err != nil {
		return AccessToken{}, err
	}
	return AccessToken{
		Token:     string(signed),
		ExpiresAt: duration,
	}, nil
}

// func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
// 	cfg := config.GetConfig()
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, jwt.ErrSignatureInvalid
// 		}
// 		return cfg.Auth.SecretKey, nil
// 	})
// 	if err != nil || !token.Valid {
// 		return nil, err
// 	}
// 	return token, nil
// }
//
// func GetUserIDFromToken(token *jwt.Token) (string, error) {
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		if userID, ok := claims["sub"].(string); ok {
// 			return userID, nil
// 		}
// 	}
// 	return "", jwt.ErrInvalidKeyType
// }

func GenerateRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	return hex.EncodeToString(token), err
}
