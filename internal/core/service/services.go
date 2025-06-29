package service

import (
	system_service "github.com/alfariiizi/go-service/internal/core/service/system"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Provide(
		system_service.NewPing,
		system_service.NewHealth,
	),
)
