package route

import (
	system_handler "github.com/alfariiizi/go-service/internal/delivery/http/route/system"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"route",
	fx.Invoke(
		system_handler.NewHandler,
	),
)
