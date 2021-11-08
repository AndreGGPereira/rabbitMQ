package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	//Pegar o valor da .env
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	if amqpServerURL == "" {
		amqpServerURL = "amqp://guest:guest@localhost:5672/"

	}
	// Cria a conexão com RabbitMQ.
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		fmt.Printf("Failed Initializing Broker Connection, error: %s", err.Error())
	}
	defer conn.Close()

	//Cria um canal de comunicação
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("erro ao tentar conectar, error: %s", err.Error())
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
		fmt.Printf("Falha ao declarar uma fila error: %s", err.Error())
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
		fmt.Printf("Falha ao consumir uma fila error: %s", err.Error())
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
			err = newClient(c)
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

//Salva um novo cliente em uma arquivo .json
func newClient(c *Cliente) error {

	var dir string
	if os.Getenv("NOVOS_CLIENTES") == "" {
		fmt.Println("warning: environment variable NOVOS_CLIENTES is not set")
		dir = "clientes"
	} else {
		dir = os.Getenv("NOVOS_CLIENTES")
	}

	//Verificar se o nome da pasta ja existe
	//Caso não exita cria com noma passado na .env
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

	//Cria o arquivo json
	newFile := filepath.Join(dir, c.UUID) + ".json"
	nf, err := os.Create(newFile)
	if err != nil {
		log.Printf(" > Falha ao criar o arquivo err: %s\n", err)
		os.Exit(3)
	}
	defer nf.Close()

	b, err := json.Marshal(c)
	if err != nil {
		log.Printf(" > Falha ao converção err: %s\n", err)
		return err
	}

	//escrevi o arquivo os dados recebido na mensagem
	if _, err := nf.Write([]byte(b)); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	return nil
}
