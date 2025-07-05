// Package core provides the core business logic of the application.
package core

import (
	"github.com/alfariiizi/go-service/internal/core/service"
	"github.com/alfariiizi/go-service/internal/core/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"core",
	usecase.Module,
	service.Module,
)
