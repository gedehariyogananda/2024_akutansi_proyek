package Models

import "time"

type MaterialStock struct {
	ID                string    `json:"id"`
	MaterialProductID string    `json:"material_product_id"`
	ProductType       string    `json:"product_type"`
	Quantity          int       `json:"quantity"`
	CurrentQuantity   int       `json:"current_quantity"`
	ExpiredDate       time.Time `json:"expired_date"`
	CreatedAt         time.Time `json:"created_at"`
}
