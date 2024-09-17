package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID           string    `json:"id"`
	CategoryName string    `json:"category_name"`
	CompanyID    string    `json:"-"`
	CreatedAt    time.Time `json:"-"`
	Company      Company   `gorm:"foreignKey:CompanyID" json:"-"`
}

// create uuid setup
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if category.ID == "" {
		category.ID = uuid.New().String()
	}

	return
}
