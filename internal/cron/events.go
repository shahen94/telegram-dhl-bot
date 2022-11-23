package cron

func (c *CronJob) Start() {
	c.scheduler.AddFunc("@every 10s", c.executor)
	c.scheduler.Start()
}

func (c *CronJob) Stop() {
	c.scheduler.Stop()
}
