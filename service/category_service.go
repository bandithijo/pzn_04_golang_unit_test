package service

import (
	"errors"
	"pzn_04_golang_unit_test/entity"
	"pzn_04_golang_unit_test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}

