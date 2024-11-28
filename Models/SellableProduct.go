package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellableProduct struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	CompanyID       string    `json:"company_id"`
	SmallesUnitID   string    `json:"smalles_unit_id"`
	CategoryID      string    `json:"category_id"`
	image           string    `json:"image"`
	Desctiption     string    `json:"description"`
	Status          bool      `json:"status"`
	HasReceipt      bool      `json:"has_receipt"`
	CurrentQuantity int       `json:"current_quantity"`
	Price           float64   `json:"price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt
}

// create uuid
func (sellableProduct *SellableProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if sellableProduct.ID == "" {
		sellableProduct.ID = uuid.New().String()
	}

	return
}
