package Dto

type SellableProductDTO struct {
	PromoID           *string `json:"promo_id"`
	Description       *string `json:"description"`
	Status            *bool   `json:"status"`
	CurrentQuantity   *int    `json:"current_quantity"`
	PrefixDeletePromo *bool   `json:"prefix_delete_promo" `
}
