package dto

type ProductDTO struct {
	Name        string `json:"productName"`
	EAN         string `json:"productEan"`
	Price       string `json:"productPrice"`
	Description string `json:"productDescription"`
	Category    string `json:"productCategory"`
}
