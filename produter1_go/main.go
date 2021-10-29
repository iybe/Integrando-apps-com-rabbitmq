package main

import (
	"fmt"
	"time"

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

	count := 1

	for {
		message := amqp.Publishing{
			Body: []byte(fmt.Sprintf("msg%d", count)),
		}

		err = channel.Publish("amq.fanout", "", false, false, message)
		if err != nil {
			panic(err)
		}

		fmt.Printf("msg%d enviada\n", count)

		count++

		time.Sleep(time.Second)
	}
}
