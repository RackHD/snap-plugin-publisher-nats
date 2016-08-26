package main

import (
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats"
)

var serverAddress, channel string

func init() {
	flag.StringVar(&serverAddress, "server-address", "nats://localhost:4222", "Address of the Nats server to listen to (default: nats://localhost:4222)")
	flag.StringVar(&channel, "channel", "Snap", "The server channel to listen to (default: Snap)")
}

func main() {
	flag.Parse()
	var err error
	var nc *nats.Conn

	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(serverAddress)
		if err == nil {
			break
		}
		log.Printf("Could not connect to Nats server. Retrying in 1s\n")
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		log.Printf("Could not connect to Nats server at %s: %s\n", serverAddress, err)
		return
	}

	log.Printf("Connected to Nats server at %s\n", serverAddress)
	log.Printf("Listening to channel %s\n", channel)

	// Simple Async Subscriber
	nc.Subscribe(channel, func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})

	runtime.Goexit()
}
