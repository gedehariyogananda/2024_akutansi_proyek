package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialProduct struct {
	ID                  string    `json:"id"`
	MaterialProductName string    `json:"material_product_name"`
	StockTypeID         string    `json:"stock_type_id"`
	IsAvailableForSale  bool      `json:"is_available_for_sale"`
	UnitPriceForSelling float64   `json:"unit_price_for_selling"`
	TotalStock          int       `json:"total_stock"`
	CompanyID           string    `json:"company_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Company             Company   `gorm:"foreignKey:CompanyID" json:"company"`
	StockType           StockType `gorm:"foreignKey:StockTypeID" json:"stock_type"`
}

// create uuid setup
func (materialProduct *MaterialProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if materialProduct.ID == "" {
		materialProduct.ID = uuid.New().String()
	}

	return
}
