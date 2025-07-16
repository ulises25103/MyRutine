package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gym-backend/database"
	"gym-backend/models"
)

type RoutineInput struct {
	Name string `json:"name"`
}

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)

	var input RoutineInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	routine := models.Routine{
		Name:   input.Name,
		UserID: userID,
	}

	if err := database.DB.Create(&routine).Error; err != nil {
		http.Error(w, "No se pudo crear rutina", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(routine)
}

func GetRoutines(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)

	var routines []models.Routine
	err := database.DB.Where("user_id = ?", userID).Find(&routines).Error
	if err != nil {
		http.Error(w, "Error al obtener rutinas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(routines)
}

func UpdateRoutine(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	params := mux.Vars(r)
	routineID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Buscar la rutina
	var routine models.Routine
	if err := database.DB.First(&routine, "id = ? AND user_id = ?", routineID, userID).Error; err != nil {
		http.Error(w, "Rutina no encontrada", http.StatusNotFound)
		return
	}

	// Leer el nuevo nombre
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	routine.Name = input.Name

	if err := database.DB.Save(&routine).Error; err != nil {
		http.Error(w, "No se pudo actualizar rutina", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(routine)
}

func DeleteRoutine(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	params := mux.Vars(r)
	routineID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var routine models.Routine
	if err := database.DB.First(&routine, "id = ? AND user_id = ?", routineID, userID).Error; err != nil {
		http.Error(w, "Rutina no encontrada", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&routine).Error; err != nil {
		http.Error(w, "No se pudo eliminar la rutina", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Rutina eliminada"})
}

func GetRoutineWithExercises(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	params := mux.Vars(r)
	routineID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Traer rutina con validación de usuario
	var routine models.Routine
	if err := database.DB.First(&routine, "id = ? AND user_id = ?", routineID, userID).Error; err != nil {
		http.Error(w, "Rutina no encontrada", http.StatusNotFound)
		return
	}

	// Traer ejercicios asociados a la rutina
	var routineExercises []models.RoutineExercise
	if err := database.DB.Where("routine_id = ?", routine.ID).Find(&routineExercises).Error; err != nil {
		http.Error(w, "Error al obtener ejercicios", http.StatusInternalServerError)
		return
	}

	// Opcional: cargar info del ejercicio (nombre y grupo muscular)
	type FullExercise struct {
		ID         uint   `json:"id"`
		Name       string `json:"name"`
		Muscle     string `json:"muscle_group"`
		Sets       int    `json:"sets"`
		Reps       int    `json:"reps"`
		RestSecs   int    `json:"rest_secs"`
	}

	var result []FullExercise
	for _, re := range routineExercises {
		var ex models.Exercise
		database.DB.First(&ex, re.ExerciseID)

		result = append(result, FullExercise{
			ID:       ex.ID,
			Name:     ex.Name,
			Muscle:   ex.MuscleGroup,
			Sets:     re.Sets,
			Reps:     re.Reps,
			RestSecs: re.RestSecs,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"routine_id": routine.ID,
		"name":       routine.Name,
		"exercises":  result,
	})
}
