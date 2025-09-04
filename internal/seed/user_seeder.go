package seed

import (
	"context"
	"log"

	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/utils"
)

type UserSeeder struct{}

func (UserSeeder) Name() string {
	return "User"
}

func (UserSeeder) Group() string {
	return "dev" // change to prod/test if needed
}

func (UserSeeder) Run(ctx context.Context, client *db.Client) error {
	err := utils.WithTx(ctx, client, func(tx *db.Tx) error {
		return nil
	})
	if err != nil {
		return err
	}

	log.Println("[Seeder] User executed successfully.")
	return nil
}
