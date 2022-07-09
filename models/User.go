package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Login    string
	Email    string
	Password string
}
