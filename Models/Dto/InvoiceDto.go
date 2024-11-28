package Dto

type PurchasedItem struct {
	ID          string  `json:"id" validate:"required"`
	Qty         int     `json:"qty" validate:"required"`
	PriceAll    float64 `json:"price_all" validate:"required"`
	PromoID     string  `json:"promo_id" validate:"omitempty"`
	PromoAmount int     `json:"promo_amount" validate:"omitempty"`
}

type InvoiceRequestDTO struct {
	CustomerName  string          `json:"customer_name" validate:"required"`
	PhoneNumber   string          `json:"phone_number" validate:"omitempty"`
	PaymentMethod string          `json:"payment_method" validate:"required"`
	Notes         string          `json:"notes" validate:"omitempty"`
	SubTotal      float64         `json:"sub_total" validate:"required"`
	InvoiceNumber string          `json:"invoice_number" validate:"required"`
	TaxID         string          `json:"tax_id" validate:"omitempty"`
	Tax           float64         `json:"tax" validate:"gte=0"`
	Status        bool            `json:"status" validate:"required"`
	Purchaseds    []PurchasedItem `json:"purchaseds" validate:"required,dive"`
}
