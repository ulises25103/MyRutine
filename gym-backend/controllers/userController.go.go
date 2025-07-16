package controllers

import (
	"encoding/json"
	"net/http"
	"gym-backend/database"
	"gym-backend/models" // Reemplazá con el path real a tu paquete models
)

// CreateUser maneja la creación de un nuevo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decodificar el JSON recibido
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Validar que los campos necesarios no estén vacíos (opcional)
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email y contraseña son requeridos", http.StatusBadRequest)
		return
	}

	// Guardar en la base de datos
	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuario creado correctamente",
	})
}
