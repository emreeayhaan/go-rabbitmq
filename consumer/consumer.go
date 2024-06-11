package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ starting winh golang in Consumer class")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMq in Consumer class")

	// opening a channel
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		"test", //queue
		"",     //consumer
		true,   //auto ack
		false,  //exclusive
		false,  //no local
		false,  //no wait
		nil,    //args
	)
	if err != nil {
		panic(err)
	}

	//consumer to producer for messages
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("waiting messages")
	<-forever
}
