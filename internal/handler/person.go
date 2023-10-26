package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/Marif226/effective-mobile-assessment/internal/lib/parsers"
	"github.com/Marif226/effective-mobile-assessment/internal/model"
	"github.com/Marif226/effective-mobile-assessment/pkg/helpers"
	"github.com/go-chi/chi/v5"
	"gopkg.in/validator.v2"
)

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	var request model.PersonCreateRequest
	err := helpers.BindRequestJSON(r, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.Validate(request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
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

func (h *Handler) listPeople(w http.ResponseWriter, r *http.Request) {
	request, err := parsers.ParsePersonListRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	response, err := h.services.PersonService.List(*request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) updatePerson(w http.ResponseWriter, r *http.Request) {
	var request model.PersonUpdateRequest
	err := helpers.BindRequestJSON(r, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.Validate(request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
    }

	response, err := h.services.PersonService.Update(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) deletePersonByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, errors.New("error: empty id").Error(), http.StatusBadRequest)
		return
	}

	err := h.services.PersonService.DeleteByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	// json.NewEncoder(w).Encode(response)
}