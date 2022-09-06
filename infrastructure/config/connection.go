package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandlers struct {
	db *gorm.DB
}

func Conection() DbHandler {
	var dsn = "host=localhost user=postgres password=dendi dbname=day7 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	dbhandler := new(DbHandlers)
	dbhandler.db = db

	return dbhandler
}

func (handler *DbHandlers) Create(value interface{}) *gorm.DB {
	return handler.db.Create(value)
}

func (handler *DbHandlers) Find(value interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(value, where...)
}

func (handler *DbHandlers) Save(value interface{}) *gorm.DB {
	return handler.db.Save(value)
}

func (handler *DbHandlers) Delete(value interface{}) *gorm.DB {
	return handler.db.Delete(value)
}
