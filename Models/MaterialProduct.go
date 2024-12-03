package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialProduct struct {
	ID                  string    `json:"id"`
	Name                string    `json:"material_product_name"`
	SmallestUnitID      string    `json:"smallest_unit_id"`
	UnitPriceForSelling float64   `json:"unit_price_for_selling"`
	CurrentQuantity     int       `json:"current_quantity"`
	Sku                 string    `json:"sku"`
	CompanyID           string    `json:"company_id"`
	CategoryID          string    `json:"category_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           gorm.DeletedAt
}

// create uuid setup
func (materialProduct *MaterialProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if materialProduct.ID == "" {
		materialProduct.ID = uuid.New().String()
	}

	return
}
