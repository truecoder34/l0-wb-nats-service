package models

import (
	uuid "github.com/satori/go.uuid"
)

type Payment struct {
	Entity
	Transaction Transaction `gorm:"foreignKey:OrderUID" json:"transaction"`
	//TransactionID uuid.UUID   `gorm:"type:uuid;column:transaction;not null" json:"transaction"`
	RequestID    uuid.UUID `gorm:"type:uuid;" json:"request_id"`
	Currency     string    `gorm:"size:255;not null;" json:"currency"`
	Provider     string    `gorm:"size:255;not null;" json:"provider"`
	Amount       string    `gorm:"size:255;not null;" json:"amount"`
	PaymentDT    int64     `gorm:"not null;" json:"payment_dt"`
	Bank         string    `gorm:"size:255;not null;" json:"bank"`
	DeliveryCost int64     `gorm:"not null;" json:"delivery_cost"`
	GoodsTotal   int64     `gorm:"not null;" json:"goods_total"`
	CustomFee    int64     `gorm:"not null;" json:"custom_fee"`
}
