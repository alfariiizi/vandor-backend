package infrastructure

import (
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/redis"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	fx.Provide(
		db.NewDB,
		redis.NewRedis,
	),
)
