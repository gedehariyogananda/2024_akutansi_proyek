package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        string         `json:"id"`
	Name      string         `json:"name,omitempty"`
	CompanyID string         `json:"company_id,omitempty"`
	Type      string         `json:"type,omitempty"`
	Code      string         `json:"code,omitempty"`
	Status    bool           `json:"status,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// create uuid setup
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if category.ID == "" {
		category.ID = uuid.New().String()
	}

	return
}
