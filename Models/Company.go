package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	ImageCompany string    `json:"image_company"`
	CodeCompany  string    `json:"code_company"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// create uuid setup
func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if company.ID == "" {
		company.ID = uuid.New().String()
	}

	return
}
