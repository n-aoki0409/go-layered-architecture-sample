package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/n-aoki0409/go-layered-architecture-sample/domain/model"
)

func NewDB() *gorm.DB {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/sample?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
