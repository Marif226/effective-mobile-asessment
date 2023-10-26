package querybuilders

import (
	"github.com/Marif226/effective-mobile-assessment/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func BuildPersonCreateQuery(request model.Person) (string, []any, error) {
	columns := []string{"name", "surname", "age", "gender", "country"}
	values := []any{request.Name, request.Surname, request.Age, request.Gender, request.Country}

	if request.Patronymic != "" {
		columns = append(columns, "patronymic")
		values = append(values, request.Patronymic)
	}

	queryBuilder := sq.Insert(
		"people",
	).Columns(
		columns...
	).Values(
		values...
	).Suffix(
		"RETURNING *",
	).PlaceholderFormat(sq.Dollar)

	return queryBuilder.ToSql()
}

func BuildPersonListQuery(request model.PersonListRequest) (string, []any, error) {
	queryBuilder := sq.Select(
		"p.id",
		"p.name",
		"p.surname",
		"p.patronymic",
		"p.age",
		"p.gender",
		"p.country",
	).From(
		"people AS p",
	).PlaceholderFormat(sq.Dollar)

	if request.Age > 0 {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"age": request.Age},
		)
	}

	if request.Country != "" {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"country": request.Country},
		)
	}

	if request.Gender != "" {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"gender": request.Gender},
		)
	}

	return queryBuilder.ToSql()
}

func BuildPersonUpdateQuery(request model.PersonUpdateRequest) (string, []any, error) {
	setMap := make(map[string]any)

	if request.Age > 0 {
		setMap["age"] = request.Age
	}

	if request.Country != "" {
		setMap["country"] = request.Country
	}

	if request.Gender != "" {
		setMap["gender"] = request.Gender
	}

	if request.Name != "" {
		setMap["name"] = request.Name
	}

	if request.Patronymic != "" {
		setMap["patronymic"] = request.Patronymic
	}

	if request.Surname != "" {
		setMap["surname"] = request.Surname
	}

	queryBuilder := sq.Update(
		"people",
	).SetMap(
		setMap,
	).Where(
		sq.Eq{"id": request.ID},
	).Suffix(
		"RETURNING *",
	).PlaceholderFormat(sq.Dollar)

	return queryBuilder.ToSql()
}

func BuildPersonDeleteByIDQuery(id int) (string, []any, error) {
	queryBuilder := sq.Delete(
		"people",
	).Where(
		sq.Eq{"id": id},
	).PlaceholderFormat(sq.Dollar)

	return queryBuilder.ToSql()
}