package controllers

import (
	"encoding/json"
	"net/http"
	"gym-backend/database"
	"gym-backend/models"
)

type ExerciseInput struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
}

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	var input ExerciseInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	ex := models.Exercise{
		Name:        input.Name,
		MuscleGroup: input.MuscleGroup,
	}

	if err := database.DB.Create(&ex).Error; err != nil {
		http.Error(w, "No se pudo crear el ejercicio", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ex)
}

func GetExercises(w http.ResponseWriter, r *http.Request) {
	var exercises []models.Exercise
	if err := database.DB.Find(&exercises).Error; err != nil {
		http.Error(w, "Error al obtener ejercicios", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exercises)
}