package cron

import (
	"context"
	"log"
	"time"

	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
)

type Scheduler struct {
	Scheduler *gocron.Scheduler
	Client    *db.Client
}

type SchedulerParams struct {
	fx.In
	Client *db.Client
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
	}
}

func (s *Scheduler) RegisterJobs() {
	// Check if there's session last activity older than 14 days
	// s.scheduler.Every(15).Minutes().Do(func() {
	// 	log.Println("[cron] Revoke sessions older than 14 days...")
	// 	_, err := s.revokeSessionUc.Execute(context.Background(), usecase.RevokeAllSessionsInput{
	// 		MaxLastUsed: -14,
	// 	})
	// 	if err != nil {
	// 		log.Printf("[cron] Error revoke old sessions: %v", err)
	// 		return
	// 	}
	// })
}
