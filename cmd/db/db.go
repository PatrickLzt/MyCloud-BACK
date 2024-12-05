package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {

	dbURL := "postgres://postgres:21622292a@localhost:5432/db_test?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &File{})

	return db
}
