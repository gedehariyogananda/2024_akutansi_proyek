package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialProduct struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CompanyID       string `json:"company_id"`
	SmallestUnitID  string `json:"smallest_unit_id"`
	CategoryID      string `json:"category_id"`
	Status          bool   `json:"status"`
	CurrentQuantity int    `json:"current_quantity"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       gorm.DeletedAt
	Unit            Unit `json:"unit" gorm:"foreignKey:SmallestUnitID"`
}

// create uuid setup
func (materialProduct *MaterialProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if materialProduct.ID == "" {
		materialProduct.ID = uuid.New().String()
	}

	return
}
