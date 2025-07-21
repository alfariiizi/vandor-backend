package main

import (
	"fmt"
	"log"

	"github.com/alfariiizi/vandor/cmd/service/utils"
)

func main() {
	err := utils.RegenerateServicesGo()
	if err != nil {
		log.Fatalf("❌ Failed to update usecases.go: %v", err)
	}
	err = utils.RegenerateGroupServiceGo()
	if err != nil {
		log.Fatalf("❌ Failed to update group service.go: %v", err)
	}

	fmt.Printf("✅ Successfully regenerated service")
}
