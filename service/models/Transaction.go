package models

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Entity

	OrderUID          string `gorm:"size:255;not null;unique" json:"order_uid"`
	TrackNumber       string `gorm:"size:255;not null;unique" json:"track_number"`
	Entry             string `gorm:"size:255;not null;" json:"entry"`
	Locale            string `gorm:"size:255;not null;" json:"locale"`
	InternalSignature string `gorm:"size:255;" json:"internal_signature"`
	CustomerID        string `gorm:"size:255;" json:"customer_id"`
	DeliveryService   string `gorm:"size:255;" json:"delivery_service"`
	Shardkey          string `gorm:"size:255;" json:"shardkey"`
	SmID              int64  `gorm:"not null;" json:"sm_id"`
	DateCreated       string `gorm:"size:255;" json:"date_created"`
	OofShard          string `gorm:"size:255;" json:"oof_shard"`

	Items    []Item   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Delivery Delivery `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment  Payment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //;references:transaction"`
}

/*
	Find all Transactions entity
*/
func (trn *Transaction) FindAllTransactions(db *gorm.DB) (*[]Transaction, error) {
	var err error
	transactions := []Transaction{}

	// WORKS. BUT FAILS WITH 3 JOINS. WHY?!
	//err = db.Debug().Model(&Transaction{}).Limit(100).Preload("Delivery").Preload("Payment").Preload("Items").Find(&transactions).Error
	err = db.Debug().Model(&Transaction{}).Limit(100).Joins("Delivery").Joins("Payment").Preload("Items").Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	return &transactions, err
}

/*
	Find Transaction by ID
*/
func (trn *Transaction) FindTransactionByID(db *gorm.DB, tid uuid.UUID) (*Transaction, error) {
	var err error
	err = db.Debug().Model(Transaction{}).Joins("Delivery").Joins("Payment").Preload("Items").Find(&trn, "transactions.id = ?", tid).Error

	if gorm.ErrRecordNotFound == err {
		return &Transaction{}, errors.New("account entity not found in database")
	} else if err != nil {
		return &Transaction{}, err
	}
	return trn, err
}

/*
	Create nested transaction received from NATS STREAMING SERVER
*/
func (trn *Transaction) CreatedNestedTransaction(db *gorm.DB, transactionToCreate Transaction) (*Transaction, error) {
	var err error

	//tx := db.Debug().Model(&Transaction{}).Session(&gorm.Session{})
	tx := db.Debug().Session(&gorm.Session{})
	err = tx.Model(&Transaction{}).Create(&trn).Error
	if err != nil {
		return &Transaction{}, err
	}

	return trn, nil
}
