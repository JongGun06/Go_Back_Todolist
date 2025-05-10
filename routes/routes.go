package routes

import (
	"github.com/gorilla/mux"
	"todolist_back_one/handlers"
)

// SetupRoutes настраивает маршруты
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE", "OPTIONS")
}