package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WaitingList struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Company      string    `json:"company"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	BusinessType string    `json:"business_type"`
	CreatedAt    time.Time `json:"created_at"`
}

func (waitingList *WaitingList) BeforeCreate(tx *gorm.DB) (err error) {
	if waitingList.ID == "" {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}

		waitingList.ID = uuidV7.String()
	}

	return
}
