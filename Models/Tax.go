package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tax struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Precentage int64          `json:"precentage"`
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Tax) TableName() string {
	return "tax"
}

func (u *Tax) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		u.ID = uuid.String()
	}

	return

}
