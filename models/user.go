package models

type User struct {
	GormModel
	Email       string    `json:"email"`
	UID         string    `json:"uid"`
	IsSuperUser bool      `json:"isSuperUser"`
	Cart        Cart      `json:"cart"`
	Orders      []Order   `json:"orders"`
	Addresses   []Address `json:"addresses"`
}

type Address struct {
	GormModel
	UserID      uint   `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalode"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber"`
}
