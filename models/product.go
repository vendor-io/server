package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	EAN         uint8
	Price       uint
	Description string
	CategoryID  uint
	Category    Category
	LeftInStock int
}

type ProductInOrder struct {
	ProductID uint
	OrderID   uint
	Product   Product
	Name      string
	Price     uint
	Amount    uint
}
