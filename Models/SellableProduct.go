package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellableProduct struct {
	ID              string         `json:"id"`
	Name            string         `json:"name,omitempty"`
	CompanyID       string         `json:"company_id,omitempty"`
	SmallestUnitID  string         `json:"smallest_unit_id,omitempty"`
	CategoryID      string         `json:"category_id,omitempty"`
	image           string         `json:"image,omitempty"`
	Desctiption     string         `json:"description,omitempty"`
	Status          bool           `json:"status,omitempty"`
	HasReceipt      bool           `json:"has_receipt,omitempty"`
	CurrentQuantity int            `json:"current_quantity,omitempty"`
	Price           float64        `json:"price,omitempty"`
	CreatedAt       *time.Time     `json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
	Unit            *Unit          `json:"unit,omitempty" gorm:"foreignKey:SmallestUnitID;references:ID"`
}

// create uuid
func (sellableProduct *SellableProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if sellableProduct.ID == "" {
		sellableProduct.ID = uuid.New().String()
	}

	return
}
