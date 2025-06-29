package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"time"
)

type TokenCodePayload struct {
	SessionID string    `json:"session_id"`
	IssuedAt  time.Time `json:"issued_at"`
}

func EncryptedTokenCode(sessionID string, secretKey []byte) (string, error) {
	payload := TokenCodePayload{
		SessionID: sessionID,
		IssuedAt:  time.Now(),
	}

	// Serialize to JSON
	plainText, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create AES cipher
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)

	// Return base64 string
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func DecryptTokenCode(encoded string, secretKey []byte) (*TokenCodePayload, error) {
	cipherText, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, err
	}

	nonce, cipherData := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}

	var payload TokenCodePayload
	if err := json.Unmarshal(plainText, &payload); err != nil {
		return nil, err
	}

	// Optionally validate expiry
	// if time.Now().After(payload.Expiry) {
	//     return nil, fmt.Errorf("code expired")
	// }

	return &payload, nil
}
