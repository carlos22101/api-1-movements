package applications

import (
	"encoding/json"
	"os"
	"recibe/src/movement/domain/entities"
	"recibe/src/movement/domain/repositories"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MovementUseCase struct {
	Repo       repositories.MovementRepository
	RabbitChan *amqp.Channel
}

func NewMovementUseCase(repo repositories.MovementRepository, ch *amqp.Channel) *MovementUseCase {
	return &MovementUseCase{
		Repo:       repo,
		RabbitChan: ch,
	}
}

func (uc *MovementUseCase) CreateMovement(m *entities.Movement) error {
	// 1. Guardar en BD
	err := uc.Repo.Create(m)
	if err != nil {
		return err
	}

	// 2. Publicar en RabbitMQ
	queueName := os.Getenv("RABBITMQ_QUEUE")
	body, _ := json.Marshal(m) // convertir a JSON
	err = uc.RabbitChan.Publish(
		"",        // exchange
		queueName, // routing key (nombre de la cola)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}
