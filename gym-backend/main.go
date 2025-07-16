package main

import (
	"gym-backend/database"
	"gym-backend/models"
	"log"
	"net/http"
	"gym-backend/routes"
)


func main() {
	database.Connect()

	// Migrar modelos
	database.DB.AutoMigrate(
		&models.User{},
		&models.Exercise{},
		&models.Routine{},
		&models.RoutineExercise{},
	)

	router := routes.SetupRoutes()

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}