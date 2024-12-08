package models

// Product represents a product entity.
type Product struct {
	ID                 string   `json:"product_id"`
	UserID             string   `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	ProductPrice       float64  `json:"product_price"`
}
