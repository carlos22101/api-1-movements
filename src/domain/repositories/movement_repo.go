package repositories

import "recibe/src/domain/entities"

type MovementRepository interface {
    Create(m *entities.Movement) error
    GetMovements() ([]entities.Movement, error)
}