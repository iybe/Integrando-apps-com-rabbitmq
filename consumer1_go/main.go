package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	url := "amqp://admin:admin@localhost:5672"

	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare("consumer1_go", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind("consumer1_go", "", "amq.fanout", false, nil)
	if err != nil {
		return
	}

	msgs, err := channel.Consume("consumer1_go", "amq.fanout", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		fmt.Printf("Received a message: %s\n", d.Body)
	}
}
