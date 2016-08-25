package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/nats-io/nats"
)

var serverAddress, channel string

func init() {
	flag.StringVar(&serverAddress, "server-address", "nats://localhost:4222", "Address of the Nats server to listen to (default: nats://localhost:4222)")
	flag.StringVar(&channel, "channel", "Snap", "The server channel to listen to (default: Snap)")
}

func main() {
	flag.Parse()

	nc, err := nats.Connect(serverAddress)
	if err != nil {
		log.Printf("Could not connect to Nats server: %s\n", err)
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
