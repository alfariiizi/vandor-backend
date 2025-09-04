package main

import (
	"context"
	"os"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/seed"
)

func main() {
	config.GetConfig()
	// group, _ := cmd.Flags().GetString("group")
	args := os.Args

	ctx := context.Background()
	client := db.NewDB() // load via Fx/config

	if len(args) > 1 {
		seed.RunOne(ctx, client, args[1])
		return
	}

	seed.RunAll(ctx, client, "dev")
}
