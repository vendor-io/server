package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	Slug        string
	ItemsAmount uint
}
