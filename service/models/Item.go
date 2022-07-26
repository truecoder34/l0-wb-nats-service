package models

import uuid "github.com/satori/go.uuid"

type Item struct {
	Entity
	ChrtID      int64  `gorm:"not null;" json:"chrt_id"`
	TrackNumber string `gorm:"size:255;not null;" json:"track_number"`
	Price       int64  `gorm:"not null;" json:"price"`
	RID         string `gorm:"size:255;not null;" json:"rid"`
	Name        string `gorm:"size:255;not null;" json:"name"`
	Sale        int64  `gorm:"not null;" json:"sale"`
	Size        string `gorm:"size:255;not null;" json:"size"`
	TotalPrice  int64  `gorm:"not null;" json:"total_price"`
	NmID        int64  `gorm:"not null;" json:"nm_id"`
	Brand       string `gorm:"size:255;not null;" json:"brand"`
	Status      int64  `gorm:"not null;" json:"status"`

	TransactionID uuid.UUID
}
