package dto

import "keyboardify-server/models"

type CartDTO struct {
	Products []models.Product
}

type CartWithTotalPriceDTO struct {
	Products   []models.Product
	TotalPrice uint64
}

type CartProductDTO struct {
	ProductID uint   `json:"productId"`
	UserID    string `json:"userId"`
}
