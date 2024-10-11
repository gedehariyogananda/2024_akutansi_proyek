package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleableProductToping struct {
	ID                string          `json:"id"`
	SaleableProductID string          `json:"saleable_product_id"`
	TopingID          string          `json:"toping_id"`
	CompanyID         string          `json:"company_id"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
	SaleableProduct   SaleableProduct `gorm:"foreignKey:SaleableProductID"`
	Toping            Toping          `gorm:"foreignKey:TopingID"`
}

// create uuid
func (saleableProductToping *SaleableProductToping) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if saleableProductToping.ID == "" {
		saleableProductToping.ID = uuid.New().String()
	}

	return
}
