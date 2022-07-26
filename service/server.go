package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/truecoder34/l0-wb-nats-service/service/controllers"
	"github.com/truecoder34/l0-wb-nats-service/service/models"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//seed.Load(server.DB)
	server.DB.Debug().Migrator().DropTable("transactions", "deliveries", "items", "payments")
	server.DB.Debug().AutoMigrate(&models.Transaction{}, &models.Delivery{}, &models.Item{}, &models.Payment{})
	server.Run(":8080")

}
