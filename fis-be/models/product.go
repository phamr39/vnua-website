package models

type Product struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Handle      string     `json:"handle"`
	Description string     `json:"description"`
	ProductInfo string     `json:"product_info"`
	UserManual  string     `json:"user_manual"`
	Images      []Image    `json:"images"`
	Categories  []Category `json:"categories"`
	CreatedAt   string     `json:"created_at"`
}
