package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (m *Manager) start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to the DHL Bot, use /track to track your package, or /add to add a new package to track")
	m.bot.Send(msg)
}
