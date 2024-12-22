package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellableStock struct {
	ID                string           `json:"id"`
	SellableProductID string           `json:"sellable_product_id"`
	ProductType       string           `json:"product_type"`
	Quantity          int              `json:"quantity"`
	CurrentQuantity   int              `json:"current_quantity"`
	ExpiredDate       time.Time        `json:"expired_date"`
	CompanyID         string           `json:"company_id"`
	SellableProduct   *SellableProduct `json:"sellable_product,omitempty"`
}

func (u *SellableStock) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	return
}
