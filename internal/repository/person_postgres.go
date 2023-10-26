package repository

import (
	"database/sql"
	"reflect"
	"github.com/Marif226/effective-mobile-assessment/internal/lib/querybuilders"
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

func (r *personRepositoryImpl) Create(request model.Person) (*model.Person, error) {
	query, args, err := querybuilders.BuildPersonCreateQuery(request)
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(query, args...)
	var newPerson model.Person
	s := reflect.ValueOf(&newPerson).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	err = row.Scan(columns...)
	if err != nil {
		return nil, err
	}

	return &newPerson, nil
}

func (r *personRepositoryImpl) List(request model.PersonListRequest) ([]model.Person, error) {
	query, args, err := querybuilders.BuildPersonListQuery(request)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	people := make([]model.Person, 0)
	for rows.Next() {
		var person model.Person
		s := reflect.ValueOf(&person).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func (r *personRepositoryImpl) Update(request model.PersonUpdateRequest) (*model.Person, error) {
	query, args, err := querybuilders.BuildPersonUpdateQuery(request)
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(query, args...)
	var person model.Person
	s := reflect.ValueOf(&person).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	err = row.Scan(columns...)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *personRepositoryImpl) DeleteByID(id int) error {
	query, args, err := querybuilders.BuildPersonDeleteByIDQuery(id)
	if err != nil {
		return err
	}

	row := r.db.QueryRow(query, args...)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}