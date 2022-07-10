package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
