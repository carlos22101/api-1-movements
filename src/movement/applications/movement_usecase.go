package applications

import (
    "recibe/src/movement/domain/entities"
    "recibe/src/movement/domain/repositories"
    rabbitmqRepositories "recibe/src/services/rabbitmq/domain/repositories"
)

type MovementUseCase struct {
    Repo      repositories.MovementRepository
    Publisher rabbitmqRepositories.MovementPublisher
}

func NewMovementUseCase(repo repositories.MovementRepository, publisher rabbitmqRepositories.MovementPublisher) *MovementUseCase {
    return &MovementUseCase{
        Repo:      repo,
        Publisher: publisher,
    }
}

func (uc *MovementUseCase) CreateMovement(m *entities.Movement) error {
    // 1. Guardar en la BD
    err := uc.Repo.Create(m)
    if err != nil {
        return err
    }
    // 2. Publicar mensaje
    return uc.Publisher.Publish(m)
}
