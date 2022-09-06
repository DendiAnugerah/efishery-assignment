package repository

import (
	"efishery-assignment/infrastructure/config"
	"efishery-assignment/model"
)

type CatRepository struct {
	DbHandler config.DbHandler
}

func (db *CatRepository) Create(c model.Category) (category model.Category, err error) {
	if err = db.DbHandler.Create(&c).Error; err != nil {
		return
	}
	category = c
	return
}

func (db *CatRepository) GetById(id int) (category model.Category, err error) {
	if err = db.DbHandler.Find(&category, id).Error; err != nil {
		return
	}
	return
}

func (db *CatRepository) GetAll() (categorys model.Categories, err error) {
	if err = db.DbHandler.Find(&categorys).Error; err != nil {
		return
	}
	return
}

func (db *CatRepository) Update(c model.Category) (category model.Category, err error) {
	if err = db.DbHandler.Save(&c).Error; err != nil {
		return
	}
	category = c
	return
}

func (db *CatRepository) Delete(c model.Category) (err error) {
	if err = db.DbHandler.Delete(&c).Error; err != nil {
		return
	}
	return
}
