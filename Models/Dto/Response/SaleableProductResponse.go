package Response

import (
	"2024_akutansi_project/Models"
	"time"

	"gorm.io/gorm"
)

type SellableResponse struct {
	ID              string              `json:"id"`
	Name            *string              `json:"name,omitempty"`
	CompanyID       *string             `json:"company_id,omitempty"`
	SmallestUnitID  *string             `json:"smallest_unit_id,omitempty"`
	CategoryID      *string             `json:"category_id,omitempty"`
	Image           *string             `json:"image,omitempty"`
	Description     *string             `json:"description,omitempty"`
	Status          *bool               `json:"status,omitempty"`
	StatusDisplay   *string             `json:"status_display,omitempty"`
	HasReceipt      *bool               `json:"has_receipt,omitempty"`
	CurrentQuantity *int                `json:"current_quantity,omitempty"`
	Price           *float64            `json:"price,omitempty"`
	CreatedAt       *time.Time          `json:"created_at,omitempty"`
	UpdatedAt       *time.Time          `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt      `json:"deleted_at,omitempty"`
	Unit            *Models.Unit        `json:"unit,omitempty" gorm:"foreignKey:SmallestUnitID"`
	Category        *Models.Category    `json:"category,omitempty"`
	PromoItems      []*Models.PromoItem `json:"promo_items,omitempty"`
}
