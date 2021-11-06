package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Cliente struct {
	UUID          string `json:"uuid,omitempty"`
	Nome          string `json:"nome,omitempty"`
	Endereco      string `json:"endereco,omitempty"`
	Cadastrado_em string `json:"cadastrado_em,omitempty"`
	Atualizado_em string `json:"atualizado_em,omitempty"`
}

func main() {

	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	conn, err := amqp.Dial(amqpServerURL)
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
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
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			c := &Cliente{}
			err := json.Unmarshal(d.Body, c)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf(" > Received message: %s\n", d.Body)
			err = novoCliente(c)
			if err != nil {
				log.Printf("Error save new client, err")
			} else {
				log.Printf(" > Saved message ")
			}
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func novoCliente(c *Cliente) error {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if os.Getenv("NOVOS_CLIENTES") == "" {
		fmt.Println("warning: environment variable NOVOS_CLIENTES is not set")
	}

	dir := os.Getenv("NOVOS_CLIENTES")

	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			e := os.Mkdir(dir, 0755)
			if e != nil {
				return err
			}
		} else {
			return err
		}
	}

	newFile := filepath.Join(dir, c.UUID) + ".json"
	nf, err := os.Create(newFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer nf.Close()

	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Failed encode", err)
		return err
	}

	if _, err := nf.Write([]byte(b)); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	return nil
}
