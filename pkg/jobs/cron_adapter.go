package jobs

import "telegram-dhl/internal/cron"

func (j *Jobs) Add(executor func()) {
	cron := cron.New(executor)
	cron.Start()
	j.crons = append(j.crons, cron)
}

func (j *Jobs) Stop() {
	for _, cron := range j.crons {
		cron.Stop()
	}
}
