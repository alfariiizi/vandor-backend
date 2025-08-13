package cron

import (
	cron "github.com/alfariiizi/vandor/internal/cron/init"
	"github.com/alfariiizi/vandor/internal/cron/scheduler"
	"go.uber.org/fx"
)

var Module = fx.Module("cron",
	fx.Provide(
		cron.NewScheduler,
	),
	fx.Invoke(func(s *cron.Scheduler) {
		scheduler.RegisterJobs(s)
	}),
)
