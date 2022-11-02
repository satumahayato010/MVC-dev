package models

import "gorm.io/gorm"

type BlogEntity struct {
	gorm.Model
	Title string
	Body  string
}

func GetAll() (data []BlogEntity) {
	result := DB.Find(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data BlogEntity) {
	result := DB.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *BlogEntity) Create() {
	result := DB.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *BlogEntity) Update() {
	result := DB.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *BlogEntity) Delete() {
	result := DB.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}
