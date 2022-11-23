package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (m *Manager) track(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Tracking your package...")
	trackingCode := update.Message.CommandArguments()

	model, err := m.services.GetOrder(trackingCode)

	if err != nil {
		msg.Text = "Error tracking package"
	} else {
		msg.Text = model.LastLocation
	}
	m.bot.Send(msg)
}
