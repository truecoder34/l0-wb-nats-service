package models

type Delivery struct {
	Entity
	Transaction Transaction `gorm:"foreignKey:OrderUID" json:"transaction"`
	//TransactionID uuid.UUID   `gorm:"type:uuid;column:transaction;not null" json:"transaction"`
	Name    string `gorm:"size:255;not null;" json:"name"`
	Phone   string `gorm:"size:100;not null;unique" json:"phone"`
	Zip     string `gorm:"size:100;not null;unique" json:"zip"`
	City    string `gorm:"size:255;not null;" json:"city"`
	Address string `gorm:"size:255;not null;unique" json:"address"`
	Region  string `gorm:"size:255;not null;unique" json:"region"`
	Email   string `gorm:"size:100;not null;unique" json:"email"`
}
