package repositories

import "recibe/src/movement/domain/entities"

type MovementRepository interface {
    Create(m *entities.Movement) error
    GetMovements() ([]entities.Movement, error)
}