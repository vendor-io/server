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
}

type ProductStock struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Stock     uint
}

type ProductInCart struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Amount    uint64
	CartID    uint
}
type ProductInOrder struct {
	ProductID uint `gorm:"primaryKey"`
	Product   Product
	Amount    uint64
	OrderID   uint
}
