package models

type Customer struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	ProductInfo string `json:"product_info"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}
