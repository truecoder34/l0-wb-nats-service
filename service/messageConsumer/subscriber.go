package messageConsumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

var usageStr = `
Usage: stan-sub [options] <subject>
Options:
	-s,  --server   <url>            NATS Streaming server URL(s)
	-c,  --cluster  <cluster name>   NATS Streaming cluster name
	-id, --clientid <client ID>      NATS Streaming client ID
	-cr, --creds    <credentials>    NATS 2.0 Credentials
Subscription Options:
	--qgroup <name>                  Queue group
	--all                            Deliver all available messages
	--last                           Deliver starting with last published message
	--since  <time_ago>              Deliver messages in last interval (e.g. 1s, 1hr)
	--seq    <seqno>                 Start at seqno
	--new_only                       Only deliver new messages
	--durable <name>                 Durable subscriber name
	--unsub                          Unsubscribe the durable on exit
`

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

// NOTE: Use tls scheme for TLS, e.g. stan-sub -s tls://demo.nats.io:4443 foo
func usage() {
	log.Fatalf(usageStr)
}

func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
}

func SubscribeNats() {

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
	// Use UserCredentials
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}

	// Connect to NATS
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", URL, clusterID, clientID)

	// Process Subscriber Options.
	startOpt := stan.StartAt(pb.StartPosition_NewOnly)
	if startSeq != 0 {
		startOpt = stan.StartAtSequence(startSeq)
	} else if deliverLast {
		startOpt = stan.StartWithLastReceived()
	} else if deliverAll && !newOnly {
		startOpt = stan.DeliverAllAvailable()
	} else if startDelta != "" {
		ago, err := time.ParseDuration(startDelta)
		if err != nil {
			sc.Close()
			log.Fatal(err)
		}
		startOpt = stan.StartAtTimeDelta(ago)
	}

	subj, i := "foo", 0
	mcb := func(msg *stan.Msg) {
		i++
		printMsg(msg, i)
	}

	sub, err := sc.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	log.Printf("Listening on [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", subj, clientID, qgroup, durable)

	if showTime {
		log.SetFlags(log.LstdFlags)
	}

	//log.Printf(byte[sub))

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			if durable == "" || unsubscribe {
				sub.Unsubscribe()
			}
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone

	//sub.Unsubscribe()
}
