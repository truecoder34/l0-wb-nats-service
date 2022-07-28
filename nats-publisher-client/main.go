package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// const (
// 	port      = ":50051"
// 	clusterID = "test-cluster"
// 	clientID  = "event-store"
// )

var (
	clusterID string
	clientID  string
	URL       string
	async     bool
	userCreds string
)

type server struct{}

/*
	Publish into chanel
*/
// func publishMsg() {
// 	sc, err := stan.Connect(
// 		clusterID,
// 		clientID,
// 		stan.NatsURL(stan.DefaultNatsURL),
// 	)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// func main() {
// 	defer sc.Close()
// 	eventMsg := []byte("MESSAGE TO CHANNEL")
// 	err = sc.Publish("foo", eventMsg)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	//log.Println("Published message on channel: " + channel)
// 	log.Println("Published message on channel: ")
// }

func main() {

	URL = stan.DefaultNatsURL
	clusterID = "test-cluster"
	clientID = "stan-pub"
	async = false
	userCreds = ""

	log.SetFlags(0)
	flag.Parse()
	//args := flag.Args()

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}
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

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	// READ JSON
	// Open our jsonFile
	jsonFile, err := os.Open("model2.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened model.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	//subj, msg := args[0], []byte(args[1])
	subj, msg := "transactions", byteValue

	ch := make(chan bool)
	var glock sync.Mutex
	var guid string
	acb := func(lguid string, err error) {
		glock.Lock()
		log.Printf("Received ACK for guid %s\n", lguid)
		defer glock.Unlock()
		if err != nil {
			log.Fatalf("Error in server ack for guid %s: %v\n", lguid, err)
		}
		if lguid != guid {
			log.Fatalf("Expected a matching guid in ack callback, got %s vs %s\n", lguid, guid)
		}
		ch <- true
	}

	if !async {
		err = sc.Publish(subj, msg)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	} else {
		glock.Lock()
		guid, err = sc.PublishAsync(subj, msg, acb)
		if err != nil {
			log.Fatalf("Error during async publish: %v\n", err)
		}
		glock.Unlock()
		if guid == "" {
			log.Fatal("Expected non-empty guid to be returned.")
		}
		log.Printf("Published [%s] : '%s' [guid: %s]\n", subj, msg, guid)

		select {
		case <-ch:
			break
		case <-time.After(5 * time.Second):
			log.Fatal("timeout")
		}

	}

}
