package Response

import (
	"2024_akutansi_project/Models"
	"time"

	"gorm.io/gorm"
)

type SellableResponse struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	CompanyID       string              `json:"company_id"`
	SmallestUnitID  string              `json:"smallest_unit_id"`
	CategoryID      string              `json:"category_id"`
	Image           string              `json:"image"`
	Description     string              `json:"description"`
	Status          bool                `json:"status"`
	StatusDisplay   string              `json:"status_display"`
	HasReceipt      bool                `json:"has_receipt"`
	CurrentQuantity int                 `json:"current_quantity"`
	Price           float64             `json:"price"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
	DeletedAt       gorm.DeletedAt      `json:"deleted_at,omitempty"`
	Unit            *Models.Unit        `json:"unit,omitempty" gorm:"foreignKey:SmallestUnitID"`
	Category        *Models.Category    `json:"category,omitempty"`
	PromoItems      []*Models.PromoItem `json:"promo_items,omitempty"`
}
