package Models

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID                string  `json:"id"`
	InvoiceID         string  `json:"invoice_id"`
	SellableProductID string  `json:"sellable_product_id"`
	Quantity          int     `json:"quantity"`
	CompanyID         string  `json:"company_id"`
	Price             float64 `json:"price"`
	PromoID           *string `json:"promo_id"`
	PromoAmount       *int    `json:"promo_amount"`
}

func (item *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if item.ID == "" {
		item.ID = uuid.New().String()
		log.Printf("Generated new ID for InvoiceItem: %s", item.ID)
	}

	return
}
