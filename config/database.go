package config

import (
	"github.com/jinzhu/gorm"
	"github.com/n-aoki0409/go-layered-architecture-sample/domain/model"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(sample_db)/sample?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
