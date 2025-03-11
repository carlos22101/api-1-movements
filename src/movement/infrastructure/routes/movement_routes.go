package routes

import (
    "database/sql"
    "net/http"
    "recibe/src/movement/applications"
    "recibe/src/movement/infrastructure/controllers"
    "recibe/src/movement/infrastructure/repositories"

    amqp "github.com/rabbitmq/amqp091-go"
    "github.com/gorilla/mux"
)

func SetupMovementRoutes(router *mux.Router, db *sql.DB, rabbitChan *amqp.Channel) {

    movRepo := repositories.NewMovementMySQLRepo(db)
 
    movUseCase := applications.NewMovementUseCase(movRepo, rabbitChan)
    getMovUseCase := applications.NewGetMovementsUseCase(movRepo)

    movController := controllers.NewMovementController(movUseCase)
    getMovController := controllers.NewGetMovementsController(getMovUseCase)

    router.HandleFunc("/movements", movController.CreateMovementHandler).Methods(http.MethodPost)
    router.HandleFunc("/movements", getMovController.GetMovementsHandler).Methods(http.MethodGet)
}