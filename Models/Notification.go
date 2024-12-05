package Models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Notifications []Notification

type Notification struct {
	ID             string          `json:"id" db:"id"`
	Scheme         string          `json:"scheme" db:"scheme"`
	UserID         string          `json:"user_id" db:"user_id"`
	Status         string          `json:"status" db:"status"`
	AdditionalData json.RawMessage `json:"additional_data" db:"additional_data"`
	QueuedAt       time.Time       `json:"queued_at" db:"queued_at"`
	ScheduledAt    time.Time       `json:"scheduled_at" db:"scheduled_at"`
	SentAt         *time.Time      `json:"sent_at,omitempty" db:"sent_at"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == "" {
		n.ID = func() string {
			id, err := uuid.NewV7()
			if err != nil {
				return uuid.New().String()
			}
			return id.String()
		}()
	}

	return nil
}
