package sse

import (
	"time"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"sse",
	fx.Provide(
		func() *Manager {
			return NewManager(
				time.Second * 10,
			)
		},
	),
)
