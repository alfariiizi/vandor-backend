package main

import (
	"fmt"
	"log"

	"github.com/alfariiizi/go-service/cmd/usecase/utils"
)

func main() {
	err := utils.RegenerateUsecasesGo()
	if err != nil {
		log.Fatalf("❌ Failed to update usecases.go: %v", err)
	}

	fmt.Printf("✅ Successfully regenerated usecase")
}
