package handler

import "github.com/Marif226/effective-mobile-assessment/internal/service"

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}