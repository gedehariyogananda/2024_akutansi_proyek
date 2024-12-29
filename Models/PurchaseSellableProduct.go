package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseSellableProduct struct {
	ID                string  `json:"id" gorm:"primaryKey"`
	PurchaseID        string  `json:"purchase_id" gorm:"not null"`
	SellableProductID string  `json:"sellable_product_id" gorm:"not null"`
	UnitPrice         float64 `json:"unit_price" gorm:"not null"`
	Quantity          int     `json:"quantity" gorm:"not null"`
	CompanyID         string  `json:"company_id" gorm:"not null"`
}

func (p *PurchaseSellableProduct) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil // Explicitly return nil if no error
}
