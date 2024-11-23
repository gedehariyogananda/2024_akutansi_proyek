package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CompanyID string    `json:"company_id"`
	Type      string    `json:"type"`
	Code      string    `json:"code"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// create uuid setup
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if category.ID == "" {
		category.ID = uuid.New().String()
	}

	return
}
