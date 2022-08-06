package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	EAN         uint64
	Price       uint64
	Description string
	CategoryID  uint
	Category    Category
	MainImage   string
	Images      string
	Stock       uint
}

type ProductImages struct {
	ProductID uint `gorm:"primaryKey"`
}

type ProductInOrder struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Amount    uint64
	OrderID   uint
}
