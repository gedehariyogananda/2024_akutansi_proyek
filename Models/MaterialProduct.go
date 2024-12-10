package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialProduct struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	Sku                 string               `json:"sku"`
	CompanyID           string               `json:"company_id"`
	SmallestUnitID      string               `json:"smallest_unit_id"`
	Unit                Unit                 `json:"unit" gorm:"foreignKey:SmallestUnitID"`
	CategoryID          string               `json:"category_id"`
	Status              bool                 `json:"status"`
	CurrentQuantity     int                  `json:"current_quantity"`
	MaterialConversions []MaterialConversion `json:"material_conversion"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	DeletedAt           gorm.DeletedAt
}

// create uuid setup
func (materialProduct *MaterialProduct) BeforeCreate(tx *gorm.DB) (err error) {
	if materialProduct.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		materialProduct.ID = uuid.String()
	}

	return
}
