package Models

import "time"

type SellableStock struct {
	ID                string    `json:"id"`
	SellableProductID string    `json:"sellable_product_id"`
	ProductType       string    `json:"product_type"`
	Quantity          int       `json:"quantity"`
	CurrentQuantity   int       `json:"current_quantity"`
	ExpiredDate       time.Time `json:"expired_date"`
	CreatedAt         time.Time `json:"created_at"`
}
