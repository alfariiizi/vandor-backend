// Package core provides the core business logic of the application.
package core

import (
	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/core/usecase"
	"github.com/alfariiizi/vandor/pkg/validator"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"core",
	validator.Module,
	usecase.Module,
	service.Module,
	domain_entries.Module,
)
