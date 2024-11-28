package Models

type Receipt struct {
	ID                string `json:"id"`
	SellableProductID string `json:"sellable_product_id"`
	MaterialProductID string `json:"material_product_id"`
	Quantity          int    `json:"quantity"`
}
