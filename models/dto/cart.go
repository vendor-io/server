package dto

import "keyboardify-server/models"

type CartDTO struct {
	Products []models.Product `json:"Products"`
}

type CartWithTotalPriceDTO struct {
	Products   []ProductInCartDTO `json:"products"`
	TotalPrice uint               `json:"totalPrice"`
}

type CartProductDTO struct {
	ProductID uint   `json:"productId"`
	UserID    string `json:"userId"`
	Amount    uint   `json:"amount"`
}
