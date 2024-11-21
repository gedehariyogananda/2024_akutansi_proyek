package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Address     *string    `json:"address,omitempty"`
	Image       *string    `json:"image,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	GeographyID *string    `json:"geography_id,omitempty"`
}

// create uuid setup
func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if company.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		company.ID = uuid.String()
	}

	return
}
