package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogActivity struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Device    string    `json:"device"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (logActivity *LogActivity) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if logActivity.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		logActivity.ID = uuid.String()
	}

	return
}
