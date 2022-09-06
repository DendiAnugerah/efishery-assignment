package model

type Category struct {
	Id_category int    `json:"id_category" gorm:"primary_key"`
	Name        string `json:"name"`
}

type Categories []Category
