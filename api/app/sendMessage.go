package app

import (
	"api/app/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

//sendMessage função que envia mesagem ao RabbitMQ
func sendMessage(c *models.Cliente) error {

	//Pegar o valor da varial de ambiente
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	if amqpServerURL == "" {
		amqpServerURL = "amqp://guest:guest@localhost:5672/"

	}
	// Cria uma nova coneção com o RabbitMQ.
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		fmt.Printf("não foi possivel estabelecer a conexão com RabbitMQ, error: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("não foi possivel criar o canal de comunicação com RabbitMQ, error: %s", err.Error())
		return err
	}
	defer ch.Close()

	//Cria uma fila para que o consumer possa receber as mensagens
	q, err := ch.QueueDeclare(
		"QueueService1", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	if err != nil {
		panic("could not open RabbitMQ channel err :" + err.Error())
	}

	body, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	//Publica a mensagem
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
