package database

import "telegram-dhl/pkg/models"

func (r *Repository) Setup() {
	r.DB.AutoMigrate(&models.Order{})
}
