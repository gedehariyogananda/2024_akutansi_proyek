package Dto

import "2024_akutansi_project/Models/Common"

type PurchasedItem struct {
	ID          string   `json:"id" validate:"required"`
	Qty         int      `json:"qty" validate:"required"`
	PriceAll    float64  `json:"price_all" validate:"required"`
	PromoID     *string  `json:"promo_id" validate:"omitempty"`
	PromoAmount *float64 `json:"promo_amount" validate:"omitempty"`
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
	Status        bool            `json:"status"`
	Purchaseds    []PurchasedItem `json:"purchaseds" validate:"required,dive"`
}

type GetHistoryInvoice struct {
	Common.Query
	StartDate *string `json:"start_date"`
	EndDate   *string `json:"end_date"`
}
