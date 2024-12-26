package db

import (
	"github.com/PatrickLzt/MyCloud-BACK/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {

	dbURL := env.GetString("DB_ADDRESS", "")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &File{})

	return db
}
