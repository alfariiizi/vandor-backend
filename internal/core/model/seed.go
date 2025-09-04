package model

import (
	"context"

	"github.com/alfariiizi/vandor/internal/infrastructure/db"
)

type Seeder interface {
	Name() string
	Group() string // e.g. "dev", "prod", "test"
	Run(ctx context.Context, client *db.Client) error
}
