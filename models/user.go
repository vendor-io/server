package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string    `json:"email"`
	UID         string    `json:"uid"`
	IsSuperUser bool      `json:"isSuperUser"`
	Cart        Cart      `json:"cart"`
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
