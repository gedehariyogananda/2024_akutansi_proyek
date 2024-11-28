package Models

type InvoiceItem struct {
	ID                string  `json:"id"`
	InvoiceID         string  `json:"invoice_id"`
	SellableProductID string  `json:"sellable_product_id"`
	Quantity          int     `json:"quantity"`
	CompanyID         string  `json:"company_id"`
	Price             float64 `json:"price"`
	PromoID           *string `json:"promo_id"`
	PromoAmount       *int64  `json:"promo_amount"`
}
