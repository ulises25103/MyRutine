package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gym-backend/controllers"
	"gym-backend/middleware"
	"github.com/rs/cors"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/routines", middleware.JWTAuth(http.HandlerFunc(controllers.CreateRoutine))).Methods("POST")
	router.Handle("/routines", middleware.JWTAuth(http.HandlerFunc(controllers.GetRoutines))).Methods("GET")
	router.Handle("/routines/{id}", middleware.JWTAuth(http.HandlerFunc(controllers.UpdateRoutine))).Methods("PUT")
	router.Handle("/routines/{id}", middleware.JWTAuth(http.HandlerFunc(controllers.DeleteRoutine))).Methods("DELETE")
	router.HandleFunc("/exercises", controllers.CreateExercise).Methods("POST")
	router.HandleFunc("/exercises", controllers.GetExercises).Methods("GET")
	router.Handle("/routines/{id}/exercises", middleware.JWTAuth(http.HandlerFunc(controllers.AddExerciseToRoutine))).Methods("POST")
	router.Handle("/routines/{id}", middleware.JWTAuth(http.HandlerFunc(controllers.GetRoutineWithExercises))).Methods("GET")
	router.Handle("/routines/{id}/exercises/{eid}", middleware.JWTAuth(http.HandlerFunc(controllers.DeleteExerciseFromRoutine))).Methods("DELETE")
	router.HandleFunc("/api/register", controllers.CreateUser).Methods("POST")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API del Gestor de Rutinas"))
	})

	// Middleware de CORS
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	log.Println("Servidor corriendo en :8080")
	http.ListenAndServe(":8080", handler)

	return router
}