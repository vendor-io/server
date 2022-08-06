package models

type orderStatus string

const (
	WAITING_FOR_PAYMENT orderStatus = "Waiting for payment"
	PROCESSING          orderStatus = "Processing"
	IN_PROGRESS         orderStatus = "In progress"
	CANCELLED           orderStatus = "Cancelled"
	SENT                orderStatus = "Sent"
	FINISHED            orderStatus = "Finished"
)

type Order struct {
	GormModel
	UserID          uint             `json:"userId"`
	User            User             `json:"user"`
	ProductsInOrder []ProductInOrder `json:"productsInOrder"`
	TotalPrice      int              `json:"totalPrice"`
	OrderStatus     orderStatus      `json:"orderStatus"`
	IsPaid          bool             `json:"isPaid"`
}
