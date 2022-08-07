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
