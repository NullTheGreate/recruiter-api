package routes

import (
	"recruiter/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router, h *handlers.UserHandler) {
	r.HandleFunc("/api/users", h.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/login", h.UserLogin).Methods("POST")
}
