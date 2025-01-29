package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialStock struct {
	ID                string          `json:"id"`
	MaterialProductID string          `json:"material_product_id"`
	ProductType       string          `json:"product_type"`
	Quantity          int             `json:"quantity"`
	CurrentQuantity   int             `json:"current_quantity"`
	ExpiredDate       time.Time       `json:"expired_date"`
	CompanyID         string          `json:"company_id"`
	MaterialProduct   MaterialProduct `json:"material_product"`
	CreatedAt         *time.Time      `json:"created_at"`
}

func (u *MaterialStock) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	if u.CreatedAt == nil {
		now := time.Now()
		u.CreatedAt = &now
	}

	return
}
