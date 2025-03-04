package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Promo struct {
	ID         string          `json:"id"`
	Name       string          `json:"name,omitempty"`
	Type       string          `json:"type,omitempty"`
	StartDate  string          `json:"start_date,omitempty"`
	EndDate    string          `json:"end_date,omitempty"`
	IsAll      bool            `json:"is_all,omitempty"`
	CompanyID  string          `json:"company_id,omitempty"`
	Amount     float64         `json:"amount,omitempty"`
	PromoItems []*PromoItem    `json:"promo_items,omitempty"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// create uuid
func (promo *Promo) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if promo.ID == "" {
		promo.ID = uuid.New().String()
	}

	return
}
