package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func New() *Repository {
	dsn := "host=localhost user=telegram_dhl password=telegram_dhl dbname=telegram_dhl port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return &Repository{
		DB: db,
	}
}
