package routes

import (
	"recruiter/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupApplicantRoutes(r *mux.Router, h *handlers.ApplicantHandler) {
	r.HandleFunc("/api/applicants", h.GetApplicants).Methods("GET")
	r.HandleFunc("/api/applicants", h.CreateApplicant).Methods("POST")
	r.HandleFunc("/api/applicants/{id}", h.GetApplicant).Methods("GET")
	r.HandleFunc("/api/applicants/{id}", h.UpdateApplicant).Methods("PUT")
	r.HandleFunc("/api/applicants/{id}", h.DeleteApplicant).Methods("DELETE")
}
