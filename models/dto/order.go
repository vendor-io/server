package dto

import (
	"keyboardify-server/models"
	"time"
)

type OrderDTO struct {
	UserID          string             `json:"userId"`
	AddressID       uint               `json:"addressId"`
	ProductsInOrder []ProductInCartDTO `json:"productsInOrder"`
	TotalPrice      uint               `json:"totalPrice"`
}

type OrderItemDTO struct {
	ID              uint                    `json:"id"`
	CreatedAt       time.Time               `json:"createdAt"`
	UserID          uint                    `json:"userId"`
	UID             string                  `json:"uid"`
	Address         PlainAddressDTO         `json:"address"`
	ProductsInOrder []models.ProductInOrder `json:"productsInOrder"`
	TotalPrice      uint                    `json:"totalPrice"`
	OrderStatus     string                  `json:"orderStatus"`
	IsPaid          bool                    `json:"isPaid"`
}
