package handlers

import (
	"telegram-dhl/pkg/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func serializeOrders(orders []models.Order) string {
	var message string
	for _, order := range orders {
		message += order.TrackingNumber + " " + order.LastLocation + "\n"
	}
	return message
}

func (m *Manager) orders(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Checking...")

	message := ""

	orders, err := m.services.GetOrders(update.Message.Chat.ID)

	if err != nil {
		message = "Error while fetching packages"
	} else {
		message = serializeOrders(orders)
	}
	msg.Text = message
	m.bot.Send(msg)
}
