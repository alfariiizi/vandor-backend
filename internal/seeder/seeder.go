package seeder

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/config"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/alfariiizi/vandor/internal/utils"
)

func GenerateAdmin(db *db.Client) error {
	cfg := config.GetConfig()

	tx, err := db.Tx(context.Background())
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	password, err := utils.HashPassword(cfg.Superadmin.Password)
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	_, err = tx.User.Create().
		SetEmail(cfg.Superadmin.Email).
		SetFirstName("Super").
		SetLastName("Admin").
		SetRole(user.RoleSUPERADMIN).
		SetPasswordHash(*password).
		Save(context.Background())
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Seeder Super Admin created successfully")

	return nil
}
