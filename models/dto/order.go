package dto

type OrderDTO struct {
	UserID          string             `json:"userId"`
	AddressID       uint               `json:"addressId"`
	ProductsInOrder []ProductInCartDTO `json:"productsInOrder"`
	TotalPrice      uint               `json:"totalPrice"`
}
