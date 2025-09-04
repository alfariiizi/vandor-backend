package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alfariiizi/vandor/cmd/utils"
)

func main() {
	name := os.Args[1]
	fileName := strings.ToLower(name) + "_seeder.go"
	path := filepath.Join("internal", "seed", fileName)

	template := fmt.Sprintf(`package seed

import (
	"context"
	"log"

	"%s/internal/infrastructure/db"
)

type %sSeeder struct{}

func (%sSeeder) Name() string {
	return "%s"
}

func (%sSeeder) Group() string {
	return "dev" // change to prod/test if needed
}

func (%sSeeder) Run(ctx context.Context, client *db.Client) error {
	// TODO: implement seeding logic here
	log.Println("[Seeder] %s executed successfully.")
	return nil
}
`, utils.GetModuleName(), name, name, name, name, name, name)

	err := os.WriteFile(path, []byte(template), 0644)
	if err != nil {
		fmt.Printf("[ERROR] Failed to create seed file: %s\n", err)
		os.Exit(1)
	}
}
