package messageConsumer

import (
	"log"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// const (
// 	clusterID   string = "test-cluster"
// 	clientID    string = "l0ID"
// 	URL         string = stan.DefaultNatsURL
// 	userCreds   string = ""
// 	showTime    bool   = false
// 	qgroup      string = ""
// 	unsubscribe bool   = true
// 	startSeq    uint64 = 0
// 	startDelta  string = ""
// 	deliverAll  bool   = true
// 	newOnly     bool   = false
// 	deliverLast bool   = false
// 	durable     string = "my-durable"
// )

func SubscribeSimpleNats() {
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
		log.Print("Got new transaction! ", m.Subject)
		log.Print("Transaction data : ", string(m.Data))

		i++
		printMsg(m, i)
	}

	conn.QueueSubscribe("transactions", qgroup, msgHandle, stan.DurableName(durable))
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}

	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", URL, clusterID, clientID)
	defer conn.Close()
}
