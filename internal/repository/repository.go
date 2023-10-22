package repository

import (
	"database/sql"

	"github.com/Marif226/effective-mobile-assessment/internal/model"
)

type PersonRepository interface {
	Create(request model.PersonCreateRequest) (*model.Person, error)
	List(request model.PersonListRequest) ([]model.Person, error)
	Update(request model.PersonUpdateRequest) (*model.Person, error)
	DeleteByID(id int) error
}

type Repository struct {
	PersonRepository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		PersonRepository: NewPersonPostgresRepository(db),
	}
}