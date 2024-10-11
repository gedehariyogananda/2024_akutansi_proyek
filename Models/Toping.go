package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Toping struct {
	ID          string    `json:"id"`
	TopingName  string    `json:"toping_name"`
	CompanyID   string    `json:"-"`
	PriceToping float64   `json:"price_toping"`
	CreatedAt   time.Time `json:"created_at"`
	Company     Company   `gorm:"foreignKey:CompanyID"`
}

// create uuid
func (toping *Toping) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if toping.ID == "" {
		toping.ID = uuid.New().String()
	}

	return
}
