package Models

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID                string   `json:"id,omitempty"`
	InvoiceID         string   `json:"invoice_id,omitempty"`
	SellableProductID string   `json:"sellable_product_id,omitempty"`
	Quantity          int      `json:"quantity,omitempty"`
	CompanyID         string   `json:"company_id,omitempty"`
	Price             float64  `json:"price,omitempty"`
	PromoID           *string  `json:"promo_id,omitempty"`
	PromoAmount       *int     `json:"promo_amount,omitempty"`
	Total             int64    `json:"total,omitempty"`
	Invoice           *Invoice `json:"invoice,omitempty"`
}

func (item *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if item.ID == "" {
		item.ID = uuid.New().String()
		log.Printf("Generated new ID for InvoiceItem: %s", item.ID)
	}

	return
}
