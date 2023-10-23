package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Marif226/effective-mobile-assessment/internal/model"
	"github.com/Marif226/effective-mobile-assessment/pkg/helpers"
)

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	var request model.PersonCreateRequest
	err := helpers.BindRequestJSON(r, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Name == "" || request.Surname == "" {
		http.Error(w, errors.New("error: empty fields").Error(), http.StatusBadRequest)
		return
	}

	response, err := h.services.PersonService.Create(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}