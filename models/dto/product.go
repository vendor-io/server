package dto

type ProductDTO struct {
	Name        string `json:"name"`
	EAN         string `json:"ean"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type ProductInCartDTO struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	MainImage    string `json:"mainImage"`
	Price        uint   `json:"price"`
	Amount       uint   `json:"amount"`
	TotalPrice   uint   `json:"totalPrice"`
	CategoryName string `json:"categoryName"`
	CategorySlug string `json:"categorySlug"`
}
