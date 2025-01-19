package repository

import "pzn_04_golang_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
