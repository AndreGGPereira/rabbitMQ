package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func sendMessage(c *Cliente) error {

	//url := os.Getenv("AMQP_URL")
	//	if url == "" {
	//		url = "amqp://guest:guest@localhost:5672"
	//	}

	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	// Define RabbitMQ server URL.

	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	//	conn, err := amqp.Dial(url)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"QueueService1", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	body, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		return err
	}
	return nil
}
