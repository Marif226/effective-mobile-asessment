package service

import (
	"github.com/Marif226/effective-mobile-assessment/internal/model"
	"github.com/Marif226/effective-mobile-assessment/internal/repository"
)

type PersonService interface {
	Create(request model.PersonCreateRequest) (*model.Person, error)
	List(request model.PersonListRequest) ([]model.Person, error)
	Update(request model.PersonUpdateRequest) (*model.Person, error)
	DeleteByID(id int) error
}

type Service struct {
	PersonService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		PersonService: NewPersonService(repo),
	}
}