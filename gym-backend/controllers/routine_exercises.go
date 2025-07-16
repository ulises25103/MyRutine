package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gym-backend/database"
	"gym-backend/models"
)

type RoutineExerciseInput struct {
	ExerciseID uint `json:"exercise_id"`
	Sets       int  `json:"sets"`
	Reps       int  `json:"reps"`
	RestSecs   int  `json:"rest_secs"`
}

func AddExerciseToRoutine(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	params := mux.Vars(r)
	routineID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de rutina inválido", http.StatusBadRequest)
		return
	}

	// Validar que la rutina sea del usuario
	var routine models.Routine
	if err := database.DB.First(&routine, "id = ? AND user_id = ?", routineID, userID).Error; err != nil {
		http.Error(w, "Rutina no encontrada", http.StatusNotFound)
		return
	}

	var input RoutineExerciseInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	re := models.RoutineExercise{
		RoutineID:  routine.ID,
		ExerciseID: input.ExerciseID,
		Sets:       input.Sets,
		Reps:       input.Reps,
		RestSecs:   input.RestSecs,
	}

	if err := database.DB.Create(&re).Error; err != nil {
		http.Error(w, "Error al asignar ejercicio", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(re)
}

func DeleteExerciseFromRoutine(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	params := mux.Vars(r)

	routineID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de rutina inválido", http.StatusBadRequest)
		return
	}

	exerciseID, err := strconv.Atoi(params["eid"])
	if err != nil {
		http.Error(w, "ID de ejercicio inválido", http.StatusBadRequest)
		return
	}

	// Verificar que la rutina pertenezca al usuario
	var routine models.Routine
	if err := database.DB.First(&routine, "id = ? AND user_id = ?", routineID, userID).Error; err != nil {
		http.Error(w, "Rutina no encontrada", http.StatusNotFound)
		return
	}

	// Buscar la relación rutina-ejercicio
	var rel models.RoutineExercise
	if err := database.DB.
		Where("routine_id = ? AND exercise_id = ?", routineID, exerciseID).
		First(&rel).Error; err != nil {
		http.Error(w, "Ejercicio no encontrado en la rutina", http.StatusNotFound)
		return
	}

	// Eliminar
	if err := database.DB.Delete(&rel).Error; err != nil {
		http.Error(w, "No se pudo eliminar el ejercicio", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Ejercicio eliminado de la rutina"})
}
