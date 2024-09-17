package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleableProduct struct {
	ID          string    `json:"id"`
	ProductName string    `json:"product_name"`
	UnitPrice   float64   `json:"unit_price"`
	CompanyID   string    `json:"company_id"`
	CategoryID  string    `json:"category_id"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Company     Company   `gorm:"foreignKey:CompanyID"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}

// create uuid
func (saleableProduct *SaleableProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if saleableProduct.ID == "" {
		saleableProduct.ID = uuid.New().String()
	}

	return
}
