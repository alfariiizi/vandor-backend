package infrastructure

import (
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
	"github.com/alfariiizi/go-service/internal/infrastructure/redis"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	fx.Provide(
		database.NewDB,
		redis.NewRedis,
	),
)
