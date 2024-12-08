package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialConversion struct {
	ID                string `json:"id"`
	Quantity          int    `json:"quantity"`
	MaterialProductID string `json:"material_product_id"`
	UnitID            string `json:"unit_id"`
	Unit              Unit   `json:"unit"`
	DeletedAt         gorm.DeletedAt
}

func (materialConversion *MaterialConversion) BeforeCreate(tx *gorm.DB) (err error) {
	if materialConversion.ID == "" {
		materialConversion.ID = uuid.New().String()
	}

	return
}
