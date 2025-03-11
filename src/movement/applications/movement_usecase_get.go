package applications

import (
    "recibe/src/movement/domain/entities"
    "recibe/src/movement/domain/repositories"
)

type GetMovementsUseCase struct {
    Repo repositories.MovementRepository
}

func NewGetMovementsUseCase(repo repositories.MovementRepository) *GetMovementsUseCase {
    return &GetMovementsUseCase{Repo: repo}
}

func (uc *GetMovementsUseCase) GetMovements() ([]entities.Movement, error) {
    return uc.Repo.GetMovements()
}