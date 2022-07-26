package seed

import (
	"github.com/truecoder34/l0-wb-nats-service/service/models"
	"gorm.io/gorm"
)

var transactions = []models.Transaction{
	// models.Transaction{
	// 	OrderUID:          "b563feb7b2b84b6test",
	// 	TrackNumber:       "WBILMTESTTRACK",
	// 	Entry:             "WBIL",
	// 	Locale:            "en",
	// 	InternalSignature: "",
	// 	CustomerID:        "test",
	// 	DeliveryService:   "meest",
	// 	Shardkey:          "9",
	// 	SmID:              99,
	// 	DateCreated:       "2021-11-26T06:22:19Z",
	// 	OofShard:          "1",
	// },
	models.Transaction{
		OrderUID:          "b563feb7b2b84b6test1",
		TrackNumber:       "WBILMTESTTRACK1",
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test1",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       "2022-02-26T06:22:19Z",
		OofShard:          "1",
	},
	models.Transaction{
		OrderUID:          "b563feb7b2b84b6test2",
		TrackNumber:       "WBILMTESTTRACK2",
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test2",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       "2022-01-26T06:22:19Z",
		OofShard:          "1",
	},
}

var payments = []models.Payment{
	// models.Payment{
	// 	TransactionOrderUID: "b563feb7b2b84b6test",
	// 	RequestID:           "",
	// 	Currency:            "USD",
	// 	Provider:            "wbpay",
	// 	Amount:              1817,
	// 	PaymentDT:           1637907727,
	// 	Bank:                "alpha",
	// 	DeliveryCost:        1500,
	// 	GoodsTotal:          317,
	// 	CustomFee:           0,
	// },
	models.Payment{
		TransactionOrderUID: "b563feb7b2b84b6test1",
		RequestID:           "",
		Currency:            "USD",
		Provider:            "wbpay",
		Amount:              1817,
		PaymentDT:           1637907827,
		Bank:                "alpha",
		DeliveryCost:        2000,
		GoodsTotal:          30,
		CustomFee:           2,
	},
	models.Payment{
		TransactionOrderUID: "b563feb7b2b84b6test2",
		RequestID:           "",
		Currency:            "USD",
		Provider:            "wbpay",
		Amount:              1817,
		PaymentDT:           1637906727,
		Bank:                "alpha",
		DeliveryCost:        3000,
		GoodsTotal:          11,
		CustomFee:           3,
	},
}

var deliveries = []models.Delivery{
	// models.Delivery{
	// 	Name:    "Tom Hardy",
	// 	Phone:   "+13477778888",
	// 	Zip:     "888999",
	// 	City:    "Miami",
	// 	Address: "Main Street 7",
	// 	Region:  "Florida",
	// 	Email:   "th777@gmail.com",
	// },
	models.Delivery{
		Name:    "Alex Petrov",
		Phone:   "+799913831313",
		Zip:     "105064",
		City:    "Moscow",
		Address: "Zemlyanoy Val Street 7",
		Region:  "Moscow",
		Email:   "alex.ptrv@yandex.ru",
	},
	models.Delivery{
		Name:    "John Doe",
		Phone:   "+00000000000",
		Zip:     "123456",
		City:    "Singapore",
		Address: "Unknown Street 188",
		Region:  "Singapore",
		Email:   "jd@yahoo.com",
	},
}

var items = []models.Item{
	// models.Item{
	// 	ChrtID:      9934930,
	// 	TrackNumber: "WBILMTESTTRACK",
	// 	Price:       453,
	// 	RID:         "",
	// 	Sale:        30,
	// 	Size:        "0",
	// 	TotalPrice:  317,
	// 	NmID:        2389212,
	// 	Brand:       "",
	// 	Status:      202,
	// },
	models.Item{
		ChrtID:      9934931,
		TrackNumber: "WBILMTESTTRACK2",
		Price:       556,
		RID:         "",
		Sale:        10,
		Size:        "0",
		TotalPrice:  494,
		NmID:        2389212,
		Brand:       "",
		Status:      202,
	},
	models.Item{
		ChrtID:      9934931,
		TrackNumber: "WBILMTESTTRACK3",
		Price:       1000,
		RID:         "",
		Sale:        0,
		Size:        "0",
		TotalPrice:  1000,
		NmID:        2389212,
		Brand:       "",
		Status:      202,
	},
	models.Item{},
}

func Load(db *gorm.DB) {
	if db.Migrator().HasTable(&models.Transaction{}) && db.Migrator().HasTable(&models.Delivery{}) && db.Migrator().HasTable(&models.Item{}) && db.Migrator().HasTable(&models.Payment{}) {
		db.Debug().Migrator().DropTable("transactions", "deliveries", "items", "payments")
		db.Debug().AutoMigrate(&models.Transaction{}, &models.Delivery{}, models.Item{}, models.Payment{})
	}

	for i, _ := range transactions {
		db.Debug().Model(&models.Transaction{}).Create(&transactions[i])

		deliveries[i].TransactionID = transactions[i].ID
		payments[i].TransactionID = transactions[i].ID
		items[i].TransactionID = transactions[i].ID

		db.Debug().Model(&models.Delivery{}).Create(&deliveries[i])
		db.Debug().Model(&models.Payment{}).Create(&payments[i])
		db.Debug().Model(&models.Item{}).Create(&items[i])
	}
}
