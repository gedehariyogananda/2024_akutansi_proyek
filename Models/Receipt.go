package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Receipt struct {
	ID                string           `json:"id"`
	SellableProductID string           `json:"sellable_product_id"`
	MaterialProductID string           `json:"material_product_id"`
	Quantity          int              `json:"quantity"`
	MaterialProduct   *MaterialProduct `json:"material_product"`
	SellableProduct   *SellableProduct `json:"sellable_product"`
}

func (receipt *Receipt) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if receipt.ID == "" {
		receipt.ID = uuid.New().String()
	}

	return
}

func (Receipt) TableName() string {
	return "receipt"
}
