package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gym-backend/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=seis86catorce dbname=gym_db port=25103 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("❌ Error al conectar a la base de datos: " + err.Error())
	}
	DB.AutoMigrate(&models.User{}, &models.Exercise{}, &models.Routine{}, &models.RoutineExercise{})
	fmt.Println("✅ Conexión exitosa a PostgreSQL")
}