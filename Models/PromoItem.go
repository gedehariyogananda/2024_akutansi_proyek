package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PromoItem struct {
	ID                string `json:"id"`
	PromoID           string `json:"promo_id"`
	SellableProductID string `json:"sellable_product_id"`
	Promo             *Promo `json:"promo,omitempty"`
}

// create uuid
func (promoItem *PromoItem) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if promoItem.ID == "" {
		promoItem.ID = uuid.New().String()
	}

	return
}
