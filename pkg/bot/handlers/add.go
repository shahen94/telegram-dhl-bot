package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (m *Manager) add(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Adding package...")
	trackingCode := update.Message.CommandArguments()

	err := m.services.AddOrder(trackingCode, update.Message.Chat.ID)

	if err != nil {
		msg.Text = "Error tracking package"
	} else {
		msg.Text = "Package added"
	}
	m.bot.Send(msg)
}
