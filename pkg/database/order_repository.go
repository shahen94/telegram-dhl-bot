package database

import "telegram-dhl/pkg/models"

func (r *Repository) AddOrder(trackingNumber string, chatId int64) error {
	order := models.Order{
		TrackingNumber: trackingNumber,
		ChatId:         chatId,
	}

	return r.DB.Create(&order).Error
}

func (r *Repository) GetOrder(trackingNumber string) (*models.Order, error) {
	var order models.Order

	if err := r.DB.Where("tracking_number = ?", trackingNumber).First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *Repository) GetUserPackages(chatId int64) ([]models.Order, error) {
	var orders []models.Order

	if err := r.DB.Where("chat_id = ?", chatId).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *Repository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order

	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *Repository) UpdateLastLocation(trackingNumber string, lastLocation string) error {
	return r.DB.Model(&models.Order{}).Where("tracking_number = ?", trackingNumber).Update("last_location", lastLocation).Error
}
