package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"recruiter/internal/models"
	"recruiter/internal/repository"
)

type ApplicantHandler struct {
	repo repository.ApplicantRepository
}

func NewApplicantHandler(repo repository.ApplicantRepository) *ApplicantHandler {
	return &ApplicantHandler{repo: repo}
}

// @Summary Get an applican
// @Description Get an applicant by ID
// @ID get-applicants-by-id
// @Produce json
// @Param id path string true "Applicant ID"
// @Success 200 {object} models.Applicants
// @Router /api/applicants/{id} [get]
func (h *ApplicantHandler) GetApplicants(w http.ResponseWriter, r *http.Request) {
	applicants, err := h.repo.GetApplicants()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(applicants)
}

func (h *ApplicantHandler) CreateApplicant(w http.ResponseWriter, r *http.Request) {
	var applicant models.Applicants
	err := json.NewDecoder(r.Body).Decode(&applicant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createApplicant, err := h.repo.CreateApplicant(applicant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createApplicant)
}

func (h *ApplicantHandler) GetApplicant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	applicant, err := h.repo.GetApplicant(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(applicant)
}

func (h *ApplicantHandler) UpdateApplicant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var applicant models.Applicants
	err = json.NewDecoder(r.Body).Decode(&applicant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedApplicant, err := h.repo.UpdateApplicant(id, applicant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedApplicant)
}

func (h *ApplicantHandler) DeleteApplicant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteApplicant(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
