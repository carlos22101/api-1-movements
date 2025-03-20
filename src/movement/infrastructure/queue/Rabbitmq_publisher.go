package queue

import (
    "encoding/json"
    "os"

    amqp "github.com/rabbitmq/amqp091-go"
)


type RabbitMQPublisher struct {
    Channel *amqp.Channel
}

func NewRabbitMQPublisher(ch *amqp.Channel) *RabbitMQPublisher {
    return &RabbitMQPublisher{Channel: ch}
}

func (r *RabbitMQPublisher) Publish(message interface{}) error {
    queueName := os.Getenv("RABBITMQ_QUEUE")

    body, err := json.Marshal(message)
    if err != nil {
        return err
    }

    return r.Channel.Publish(
        "",        // Exchange
        queueName, // Routing Key
        false,     // Mandatory
        false,     // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}
