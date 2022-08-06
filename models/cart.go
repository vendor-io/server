package models

type Cart struct {
	GormModel
	UserID         uint          `json:"userId"`
	ProductsInCart []CartProduct `json:"productsInCart"`
}

type CartProduct struct {
	CartID    uint `json:"cartId"`
	ProductID uint `json:"productId"`
}
