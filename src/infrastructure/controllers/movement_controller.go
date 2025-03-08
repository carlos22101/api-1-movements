package controllers

import (
	"encoding/json"
	"net/http"
	"recibe/src/applications"
	"recibe/src/domain/entities"
)

type MovementController struct {
	UseCase *applications.MovementUseCase
}

func NewMovementController(uc *applications.MovementUseCase) *MovementController {
	return &MovementController{UseCase: uc}
}

// POST /movements
func (c *MovementController) CreateMovementHandler(w http.ResponseWriter, r *http.Request) {
	var mov entities.Movement
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		http.Error(w, "Error en el body", http.StatusBadRequest)
		return
	}

	err := c.UseCase.CreateMovement(&mov)
	if err != nil {
		http.Error(w, "No se pudo crear el movimiento", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mov)
}
