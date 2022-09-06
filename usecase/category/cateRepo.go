package categoryusecase

import (
	"efishery-assignment/model"
)

type CatRepo interface {
	Create(model.Category) (model.Category, error)
	GetById(int) (model.Category, error)
	GetAll() (model.Categories, error)
	Update(model.Category) (model.Category, error)
	Delete(model.Category) error
}
