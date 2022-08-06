package models

type Category struct {
	GormModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	ItemsAmount uint   `json:"itemsAmount"`
}
