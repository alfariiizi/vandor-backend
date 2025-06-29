package cron

import (
	"go.uber.org/fx"
)

var Module = fx.Module("cron",
	fx.Provide(
		NewScheduler,
	),
	fx.Invoke(func(s *Scheduler) {
		s.RegisterJobs()
	}),
)
