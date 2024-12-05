package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type File struct {
	gorm.Model
	Title  string `json:"title"`
	Body   []byte `json:"body"`
	Status string `json:"status"`
}
