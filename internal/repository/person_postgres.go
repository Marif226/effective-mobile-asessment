package repository

import (
	"database/sql"

	"github.com/Marif226/effective-mobile-assessment/internal/model"
)

type personRepositoryImpl struct {
	db *sql.DB
}

func NewPersonPostgresRepository(db *sql.DB) *personRepositoryImpl {
	return &personRepositoryImpl{
		db: db,
	}
}

func (r personRepositoryImpl) Create(request model.Person) (*model.Person, error) {
	return nil, nil
}

func (r personRepositoryImpl) List(request model.PersonListRequest) ([]model.Person, error) {
	return nil, nil
}

func (r personRepositoryImpl) Update(request model.PersonUpdateRequest) (*model.Person, error) {
	return nil, nil
}

func (r personRepositoryImpl) DeleteByID(id int) error {
	return nil
}