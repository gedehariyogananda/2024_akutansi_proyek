package Dto

type PurchasedItem struct {
	ID          string  `json:"id" binding:"required"`
	Qty         int     `json:"qty" binding:"required"`
	PriceAll    float64 `json:"price_all"`
	PromoID     string  `json:"promo_id"`
	PromoAmount int     `json:"promo_amount"`
}

type InvoiceRequestDTO struct {
	CustomerName  string          `json:"customer_name" binding:"required"`
	PhoneNumber   string          `json:"phone_number"`
	PaymentMethod string          `json:"payment_method" binding:"required"`
	Notes         string          `json:"notes"`
	SubTotal      float64         `json:"sub_total" binding:"required"`
	InvoiceNumber string          `json:"invoice_number" binding:"required"`
	TaxID         string          `json:"tax_id"`
	Tax           float64         `json:"tax"`
	Status        bool            `json:"status" binding:"required"`
	Purchaseds    []PurchasedItem `json:"purchaseds" binding:"required,dive"`
}
