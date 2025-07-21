package main

import (
	"fmt"
	"log"

	"github.com/alfariiizi/vandor/cmd/http/utils"
)

func main() {
	err := utils.RegenerateRoutesGo()
	if err != nil {
		log.Fatalf("❌ Failed to update routes.go: %v", err)
	}

	fmt.Printf("✅ Successfully regenerated routes")
}
