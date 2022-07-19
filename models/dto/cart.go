package dto

import "keyboardify-server/models"

type CartDTO struct {
	Products   []models.ProductInCart
	TotalPrice uint64
}
