package scheduler

import (
	"log"

	cron "github.com/alfariiizi/vandor/internal/cron/init"
)

func RegisterCleanCacheJob(s *cron.Scheduler) {
	s.Scheduler.Every(1).Minute().Do(func() {
		log.Println("[cron] Running CleanCache job...")

		// payload := job.LogSystemPayload{Message: "CleanCache job executed"}
		//
		// if _, err := s.Jobs.LogSystem.Enqueue(context.Background(), payload); err != nil {
		// 	log.Printf("[cron] Error enqueueing CleanCache job: %v", err)
		// }
	})
}
