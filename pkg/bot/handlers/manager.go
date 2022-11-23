package handlers

import (
	"telegram-dhl/pkg/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Manager struct {
	bot      *tgbotapi.BotAPI
	services services.OrderServices
}

func (m *Manager) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := m.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		m.handleCommand(update)
	}
}

func NewManager(bot *tgbotapi.BotAPI, services services.OrderServices) *Manager {
	return &Manager{
		bot:      bot,
		services: services,
	}
}
