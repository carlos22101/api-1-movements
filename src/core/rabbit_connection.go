package core

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbit() (*amqp.Channel, error) {
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")

	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	// Declarar cola (asegurarnos que exista)
	queueName := os.Getenv("RABBITMQ_QUEUE")
	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
