package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID                string           `json:"id,omitempty"`
	InvoiceID         string           `json:"invoice_id,omitempty"`
	SellableProductID string           `json:"sellable_product_id,omitempty"`
	Quantity          int              `json:"quantity,omitempty"`
	CompanyID         string           `json:"company_id,omitempty"`
	Price             float64          `json:"price,omitempty"`
	PromoID           *string          `json:"promo_id,omitempty"`
	PromoAmount       *float64         `json:"promo_amount,omitempty"`
	Invoice           *Invoice         `json:"invoice,omitempty"`
	SellableProduct   *SellableProduct `json:"sellable_product,omitempty"`
	Promo             *Promo           `json:"promo,omitempty`
}

func (item *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	if item.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		item.ID = uuid.String()
	}

	return
}
