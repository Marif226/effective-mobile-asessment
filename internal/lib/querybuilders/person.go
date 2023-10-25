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