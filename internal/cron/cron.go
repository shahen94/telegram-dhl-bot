package cron

import "github.com/robfig/cron/v3"

type CronJob struct {
	scheduler *cron.Cron
	executor  func()
}

func New(executor func()) *CronJob {

	return &CronJob{
		scheduler: cron.New(),
		executor:  executor,
	}
}
