package main

import (
	"telegram-dhl/pkg/bot"
)

func main() {
	instance := bot.New()

	instance.Logger.Info("Starting bot...")

	instance.Start()
}
