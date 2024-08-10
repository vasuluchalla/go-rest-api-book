package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost:5432/bookapp?sslmode=disable")
	HandleError(err)
	return db
}

func GetDB() *gorm.DB {
	return db
}
