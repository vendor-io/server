package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID         uint
	ProductsInCart []CartProduct
}

type CartProduct struct {
	CartID    uint
	ProductID uint
}
