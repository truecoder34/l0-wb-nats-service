package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"

	"github.com/truecoder34/l0-wb-nats-service/service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.Transaction{}, &models.Payment{}, &models.Item{}, &models.Delivery{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

const (
	clusterID   string = "test-cluster"
	clientID    string = "l0ID"
	URL         string = stan.DefaultNatsURL
	userCreds   string = ""
	showTime    bool   = false
	qgroup      string = ""
	unsubscribe bool   = true
	startSeq    uint64 = 0
	startDelta  string = ""
	deliverAll  bool   = true
	newOnly     bool   = false
	deliverLast bool   = false
	durable     string = "my-durable"
)

func (server *Server) printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
	log.Printf("[#%d] Received.Data: %s\n", i, m.Data)

	tr, err := server.ParseTransactionsMessage(m.Data)
	if err != nil {
		// TODO : ADD CHECK TO PREVENT UNEXPECTED DATA PROCESSING
		log.Print(err)
	}

	log.Printf("[#%d] Received.Transaction: %s\n", i, tr)
}

func (server *Server) Run(addr string) {

	// NATS STREAMING  CONNECTION TODO: Incapsulate it into separate function
	var natsConn *nats.Conn
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer natsConn.Close()
	conn, err := stan.Connect(clusterID, clientID, stan.NatsConn(natsConn))
	if err != nil {
		log.Print(err)
	}
	//initial message handler
	i := 0
	msgHandle := func(m *stan.Msg) {
		i++
		server.printMsg(m, i)
	}
	conn.QueueSubscribe("transactions", qgroup, msgHandle, stan.DurableName(durable))
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", URL, clusterID, clientID)
	defer conn.Close()
	//messageConsumer.SubscribeSimpleNats()

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
