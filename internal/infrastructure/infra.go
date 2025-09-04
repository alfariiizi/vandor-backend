package infrastructure

import (
	"github.com/alfariiizi/vandor/internal/infrastructure/asynq"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/redis"
	"github.com/alfariiizi/vandor/internal/infrastructure/sse"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure",
	fx.Provide(
		db.NewDB,
		redis.NewRedis,
	),
	sse.Module,
	asynq.WorkerModule,
)
