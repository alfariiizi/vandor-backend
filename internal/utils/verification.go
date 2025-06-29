package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(99999) + 1     // Generate a random number between 1 and 99999
	return fmt.Sprintf("%05d", code) // Format it as a 5-digit string
}
