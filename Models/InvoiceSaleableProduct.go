package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceSaleableProduct struct {
	ID                string          `json:"id"`
	InvoiceID         string          `json:"invoice_id"`
	SaleableProductID string          `json:"saleable_product_id"`
	QuantitySold      int             `json:"quantity_sold"`
	CompanyID         string          `json:"-"`
	Notes             string          `json:"notes"`
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID" json:"-"`
	SaleableProduct   SaleableProduct `gorm:"foreignKey:SaleableProductID" json:"saleable_product"`
	Company           Company         `gorm:"foreignKey:CompanyID" json:"-"`
}

// create uuid
func (invoiceSaleableProduct *InvoiceSaleableProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if invoiceSaleableProduct.ID == "" {
		invoiceSaleableProduct.ID = uuid.New().String()
	}

	return
}
