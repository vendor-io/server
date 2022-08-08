package models

type Address struct {
	GormModel
	UserID      uint   `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber"`
}

type AddressInOrder struct {
	OrderID   uint `gorm:"primaryKey" json:"orderId"`
	AddressID uint `json:"addressId"`
}
