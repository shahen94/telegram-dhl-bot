package services

import (
	"telegram-dhl/pkg/database"
	"telegram-dhl/pkg/models"
)

type OrderServices interface {
	// AddOrder adds a new order to the database
	AddOrder(trackingNumber string, chatId int64) error
	// GetOrder returns an order by its tracking number
	GetOrder(trackingNumber string) (*models.Order, error)
	// GetOrders returns all orders for a given chat id
	GetOrders(chatId int64) ([]models.Order, error)
	// GetAll returns all orders
	GetAll() ([]models.Order, error)
	// UpdateLastLocation updates the last location of an order
	UpdateLastLocation(trackingNumber string, lastLocation string) error
}

type orderServices struct {
	Repository *database.Repository
}

func (s *orderServices) AddOrder(trackingNumber string, chatId int64) error {
	return s.Repository.AddOrder(trackingNumber, chatId)
}

func (s *orderServices) GetOrder(trackingNumber string) (*models.Order, error) {
	return s.Repository.GetOrder(trackingNumber)
}

func (s *orderServices) GetOrders(chatId int64) ([]models.Order, error) {
	return s.Repository.GetUserPackages(chatId)
}

func (s *orderServices) GetAll() ([]models.Order, error) {
	return s.Repository.GetAllOrders()
}

func (s *orderServices) UpdateLastLocation(trackingNumber string, lastLocation string) error {
	return s.Repository.UpdateLastLocation(trackingNumber, lastLocation)
}

func NewOrderServices(repository *database.Repository) OrderServices {
	return &orderServices{
		Repository: repository,
	}
}
