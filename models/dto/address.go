package dto

type AddressDTO struct {
	UserID      string `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber"`
}

type PlainAddressDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber"`
}
