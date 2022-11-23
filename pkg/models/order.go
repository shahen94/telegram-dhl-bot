package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TrackingNumber string
	LastLocation   string
	ChatId         int64
}

func (o *Order) UpdateLocation(location string) {
	o.LastLocation = location
}

func (o *Order) UpdateTrackingNumber(trackingNumber string) {
	o.TrackingNumber = trackingNumber
}

func New(trackingNumber string) *Order {
	return &Order{
		TrackingNumber: trackingNumber,
	}
}
