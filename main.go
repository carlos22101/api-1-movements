package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // Para cargar .env (go get github.com/joho/godotenv)
	"recibe/src/core"
	"recibe/src/movement/infrastructure/routes"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env (posible que no exista)")
	}

	// Conexión a la BD
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)
	}
	defer db.Close()

	// Conexión a RabbitMQ
	rabbitChan, err := core.ConnectRabbit()
	if err != nil {
		log.Fatal("Error conectando a RabbitMQ:", err)
	}
	defer rabbitChan.Close()

	// Configurar router
	router := mux.NewRouter()
	// Registrar rutas
	routes.SetupMovementRoutes(router, db, rabbitChan)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Producer API corriendo en el puerto:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
