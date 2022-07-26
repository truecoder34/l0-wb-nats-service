package models

import uuid "github.com/satori/go.uuid"

type Delivery struct {
	Entity
	//Transaction string `gorm:"size:255;not null;unique" json:"transaction"`
	//TransactionID uuid.UUID   `gorm:"type:uuid;column:transaction;not null" json:"transaction"`

	TransactionID uuid.UUID

	//TransactionID     uuid.UUID   `gorm:"type:uuid;column:transaction_id;not null" json:"transaction_id"`
	//TransactionEntity Transaction `gorm:"foreignKey:TransactionID" json:"transaction_entity"`

	Name    string `gorm:"size:255;not null;" json:"name"`
	Phone   string `gorm:"size:100;not null;unique" json:"phone"`
	Zip     string `gorm:"size:100;not null;unique" json:"zip"`
	City    string `gorm:"size:255;not null;" json:"city"`
	Address string `gorm:"size:255;not null;unique" json:"address"`
	Region  string `gorm:"size:255;not null;unique" json:"region"`
	Email   string `gorm:"size:100;not null;unique" json:"email"`
}
