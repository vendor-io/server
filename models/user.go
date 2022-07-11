package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login       string    `json:"login"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	ItemsInCart []Product `json:"itemsInCart" gorm:"many2many:user_carts;"`
	Orders      []Order   `json:"orders"`
	Addresses   []Address `json:"addresses"`
}

type Address struct {
	gorm.Model
	UserID      uint
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalode"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber"`
}
