package models

type Product struct {
	GormModel
	Name        string   `json:"name"`
	EAN         uint64   `json:"ean"`
	Price       uint     `json:"price"`
	Description string   `json:"description"`
	CategoryID  uint     `json:"categoryId"`
	Category    Category `json:"category"`
	MainImage   string   `json:"mainImage"`
	Images      string   `json:"images"`
	Stock       uint     `json:"stock"`
}

type ProductInOrder struct {
	ProductID uint    `gorm:"primaryKey" json:"productId"`
	Product   Product `json:"product"`
	Amount    uint64  `json:"amount"`
	OrderID   uint    `json:"orderId"`
}
