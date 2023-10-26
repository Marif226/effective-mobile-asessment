package parsers

import (
	"net/http"
	"strconv"

	"github.com/Marif226/effective-mobile-assessment/internal/model"
)

func ParsePersonListRequest(r *http.Request) (*model.PersonListRequest, error) {
	request := &model.PersonListRequest{
		Country: r.URL.Query().Get("country"),
		Gender: r.URL.Query().Get("gender"),
	}

	ageStr := r.URL.Query().Get("age")
	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			return nil, err
		}

		request.Age = age
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, err
		}

		request.Offset = offset
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, err
		}

		request.Limit = limit
	}

	return request, nil
}