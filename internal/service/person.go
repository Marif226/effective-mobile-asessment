package service

import (
	"github.com/Marif226/effective-mobile-assessment/internal/model"
	"github.com/Marif226/effective-mobile-assessment/internal/repository"
)

type personServiceImpl struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) *personServiceImpl {
	return &personServiceImpl{
		repo: repo,
	}
}

func (s personServiceImpl) Create(request model.PersonCreateRequest) (*model.Person, error) {
	newPerson := model.Person{
		Name: request.Name,
		Surname: request.Surname,
		Patronymic: request.Patronymic,
	}

	response, err := s.repo.Create(newPerson)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s personServiceImpl) List(request model.PersonListRequest) ([]model.Person, error) {
	return nil, nil
}

func (s personServiceImpl) Update(request model.PersonUpdateRequest) (*model.Person, error) {
	return nil, nil
}

func (s personServiceImpl) DeleteByID(id int) error {
	return nil
}