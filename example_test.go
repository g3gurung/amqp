package amqp_test

import (
	"context"
	"fmt"
	"log"

	"pack.ag/amqp"
)

func Example() {
	// Create client
	client, err := amqp.Dial("amqps://my-namespace.servicebus.windows.net",
		amqp.ConnSASLPlain("access-key-name", "access-key"),
	)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	// Create a receiver
	receiver, err := session.NewReceiver(
		amqp.LinkSource("/queue-name"),
		amqp.LinkCredit(10),
	)
	if err != nil {
		log.Fatal("Creating receiver link:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		// Receive next message
		msg, err := receiver.Receive(ctx)
		if err != nil {
			log.Fatal("Reading message from AMQP:", err)
		}

		// Accept message
		msg.Accept()

		fmt.Printf("Message received: %s\n", msg.Data)
	}
}
