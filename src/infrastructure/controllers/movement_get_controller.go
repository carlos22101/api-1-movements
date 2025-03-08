package controllers

import (
    "encoding/json"
    "net/http"
    "recibe/src/applications"
)

type GetMovementsController struct {
    UseCase *applications.GetMovementsUseCase
}

func NewGetMovementsController(uc *applications.GetMovementsUseCase) *GetMovementsController {
    return &GetMovementsController{UseCase: uc}
}

// GET /movements
func (c *GetMovementsController) GetMovementsHandler(w http.ResponseWriter, r *http.Request) {
    movements, err := c.UseCase.GetMovements()
    if err != nil {
        http.Error(w, "No se pudieron obtener los movimientos", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movements)
}