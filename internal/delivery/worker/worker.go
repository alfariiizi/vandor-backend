package worker

import (
	"go.uber.org/fx"
)

var Module = fx.Module("worker_delivery",
	fx.Provide(
		NewWorkerServer,
		NewWorkerClient,
		NewWorkerInspector,
	),
)
