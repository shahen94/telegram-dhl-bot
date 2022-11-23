package jobs

import (
	"telegram-dhl/internal/cron"
	"telegram-dhl/pkg/scraper"
	"telegram-dhl/pkg/services"
)

type Jobs struct {
	services services.OrderServices
	scraper  *scraper.Scraper
	crons    []*cron.CronJob
}

func New(services services.OrderServices) *Jobs {
	return &Jobs{
		services: services,
		crons:    []*cron.CronJob{},
		scraper:  scraper.New(),
	}
}
