package handler

import (
	"github.com/Marif226/effective-mobile-assessment/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/people", func(r chi.Router) {
		r.Post("/", h.createPerson)
		r.Get("/", h.listPeople)
		r.Put("/", h.updatePerson)
		r.Delete("/{id}", h.deletePersonByID)
	})

	return router
}