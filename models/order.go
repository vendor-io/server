package models

import (
	"gorm.io/gorm"
)

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
	gorm.Model
	UserID      uint
	User        User
	Products    []ProductInOrder
	TotalPrice  int
	OrderStatus orderStatus
	IsPaid      bool
}
