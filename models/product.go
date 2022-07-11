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
}

type ProductStock struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Stock     uint
}

type ProductInCart struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Amount    uint
	CartID    uint
}
type ProductInOrder struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Amount    uint
	OrderID   uint
}
