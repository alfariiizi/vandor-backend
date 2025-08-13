package croninit

import (
	"context"
	"log"
	"time"

	"github.com/alfariiizi/vandor/internal/core/job"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
)

type Scheduler struct {
	Scheduler *gocron.Scheduler
	Client    *db.Client
	Jobs      *job.Jobs
}

type SchedulerParams struct {
	fx.In
	Client *db.Client
	Jobs   *job.Jobs
}

func NewScheduler(
	lc fx.Lifecycle,
	params SchedulerParams,
) *Scheduler {
	s := gocron.NewScheduler(time.Local)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("[cron] Starting scheduler...")
			s.StartAsync() // Run in background
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("[cron] Stopping scheduler...")
			s.Stop()
			return nil
		},
	})

	return &Scheduler{
		Scheduler: s,
		Client:    params.Client,
		Jobs:      params.Jobs,
	}
}
