package models

type Store struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address"`
	BusinessModel string `json:"business_model"`
	Message       string `json:"message"`
}
