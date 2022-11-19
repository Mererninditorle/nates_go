package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	go sub()

	for {

	}
}

func  sub() {
	// Connect to a server
	nc, _ := nats.Connect("nats://95.165.107.100:4222")

	// Simple Publisher
	// nc.Publish("test_shapovalov", []byte("Hello World"))

	// Responding to a request message
	nc.Subscribe("ith.class*.students", func(m *nats.Msg) {
		m.Respond([]byte("Шаповалов принял!"))
		fmt.Printf("recieved a message: %s\n", string(m.Data))
	})

	// Close connection
	// nc.Close()
}