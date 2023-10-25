package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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

func (s *personServiceImpl) Create(request model.PersonCreateRequest) (*model.Person, error) {
	var enrichInfo model.EnrichInfo
	err := enrichAPI(&enrichInfo, "https://api.agify.io/", request.Name)
	if err != nil {
		return nil, err
	}
	err = enrichAPI(&enrichInfo, "https://api.genderize.io/", request.Name)
	if err != nil {
		return nil, err
	}

	err = enrichAPI(&enrichInfo, "https://api.nationalize.io/", request.Name)
	if err != nil {
		return nil, err
	}

	maxProb := 0.0
	country := ""
	for _, c := range enrichInfo.Country {
		if c.Probability > maxProb {
			maxProb = c.Probability
			country = c.CountryID
		}
	}

	newPerson := model.Person{
		Name: request.Name,
		Surname: request.Surname,
		Patronymic: request.Patronymic,
		Age: enrichInfo.Age,
		Gender: enrichInfo.Gender,
		Country: country,
	}

	response, err := s.repo.Create(newPerson)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *personServiceImpl) List(request model.PersonListRequest) ([]model.Person, error) {
	return nil, nil
}

func (s *personServiceImpl) Update(request model.PersonUpdateRequest) (*model.Person, error) {
	return nil, nil
}

func (s *personServiceImpl) DeleteByID(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = s.repo.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}

func enrichAPI(enrichInfo *model.EnrichInfo, api string, name string) error {
	path := fmt.Sprintf("%s?name=%s", api, name)
	apiResp, err := http.Get(path)
	if err != nil {
		return err
	}
	defer apiResp.Body.Close()


	respBody, err := io.ReadAll(apiResp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(respBody, enrichInfo)

	return nil
}