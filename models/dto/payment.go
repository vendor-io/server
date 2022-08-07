package dto

type PaymentDTO struct {
	Amount       uint   `json:"amount"`
	Source       string `json:"source"`
	ReceiptEmail string `json:"receiptEmail"`
}

type CheckoutData struct {
	ClientSecret string `json:"clientSecret"`
}
