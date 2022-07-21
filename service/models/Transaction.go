package models

import (
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	Entity
	OrderUID    uuid.UUID `gorm:"type:uuid;primary_key;" json:"order_uid"`
	TrackNumber string    `gorm:"size:255;not null;" json:"track_number"`
	Entry       string    `gorm:"size:255;not null;" json:"entry"`
	Items       []Item
	//delivery
	//payment
	//items
	Locale            string `gorm:"size:255;not null;" json:"locale"`
	InternalSignature string `gorm:"size:255;" json:"internal_signature"`
	CustomerID        string `gorm:"size:255;" json:"customer_id"`
	DeliveryService   string `gorm:"size:255;" json:"delivery_service"`
	Shardkey          string `gorm:"size:255;" json:"shardkey"`
	SmID              string `gorm:"size:255;" json:"sm_id"`
	DateCreated       string `gorm:"size:255;" json:"date_created"`
	OofShard          string `gorm:"size:255;" json:"oof_shard"`
}
