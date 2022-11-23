package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	CommandStart  = "start"
	CommandTrack  = "track"
	CommandAdd    = "add"
	CommandOrders = "orders"
)

func (m *Manager) handleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case CommandStart:
		m.start(update)
	case CommandTrack:
		m.track(update)
	case CommandAdd:
		m.add(update)
	case CommandOrders:
		m.orders(update)
	}
}
