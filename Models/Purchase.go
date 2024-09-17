package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID                  string
	TotalPurchaseAmount float64
	CompanyID           string
	CreatedAt           time.Time
	Company             Company `gorm:"foreignKey:CompanyID"`
}

// create uuid
func (purchase *Purchase) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if purchase.ID == "" {
		purchase.ID = uuid.New().String()
	}

	return
}
