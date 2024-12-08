package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Unit struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	CompanyID        string    `json:"company_id"`
	Code             string    `json:"code"`
	Status           bool      `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        gorm.DeletedAt
}

func (u *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	return
}
